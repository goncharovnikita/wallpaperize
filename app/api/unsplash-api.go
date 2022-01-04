package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
)

var (
	unsplashAPIprefix      = "https://api.unsplash.com/"
	unsplashRandomPhotoURL = "photos/random"
)

type idGetter interface {
	GetToken() (string, error)
}

// UnsplashAPI implementation
type UnsplashAPI struct {
	idGetter idGetter
}

func NewUnsplashAPI(idGetter idGetter) *UnsplashAPI {
	return &UnsplashAPI{
		idGetter: idGetter,
	}
}

type unsplashRandomImageURLs struct {
	RAW string `json:"raw"`
}

type UnsplashRandomImage struct {
	ID          string                  `json:"id"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
	Width       int                     `json:"width"`
	Height      int                     `json:"height"`
	Description string                  `json:"description"`
	URLs        unsplashRandomImageURLs `json:"urls"`
}

type UnsplashRandomImageResponse struct {
	Data               UnsplashRandomImage
	RateLimitRemaining int
}

func (u *UnsplashAPI) GetRandomImage() (*UnsplashRandomImageResponse, error) {
	token, err := u.idGetter.GetToken()
	if err != nil {
		return nil, err
	}

	var client http.Client

	url := unsplashAPIprefix + unsplashRandomPhotoURL + "?orientation=landscape&w=1920&h=1080"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		var data interface{}
		json.NewDecoder(response.Body).Decode(&data)

		return nil, fmt.Errorf("error response status from unsplash: %d. %v", response.StatusCode, data)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	rateLimitRemainingStr := response.Header.Get("X-Ratelimit-Remaining")
	rateLimitRemaining, err := strconv.Atoi(rateLimitRemainingStr)
	if err != nil {
		return nil, err
	}

	var data UnsplashRandomImage

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &UnsplashRandomImageResponse{
		Data:               data,
		RateLimitRemaining: rateLimitRemaining,
	}, nil
}

type tokenStore interface {
	Get() ([]byte, error)
	Set([]byte) error
}

type UnsplashAuthorizer struct {
	clientID     string
	clientSecret string
	tokenStore   tokenStore
}

func NewUnsplashAuthorizer(clientID, clientSecret string, tokenStore tokenStore) *UnsplashAuthorizer {
	return &UnsplashAuthorizer{
		clientID:     clientID,
		clientSecret: clientSecret,
		tokenStore:   tokenStore,
	}
}

func (a *UnsplashAuthorizer) GetToken() (string, error) {
	currToken, err := a.tokenStore.Get()
	if err == nil && len(currToken) > 0 {
		return string(currToken), nil
	}

	authURL := url.URL{
		Scheme: "https",
		Host:   "unsplash.com",
		Path:   "/oauth/authorize",
	}
	q := authURL.Query()
	q.Set("client_id", a.clientID)
	q.Set("redirect_uri", "urn:ietf:wg:oauth:2.0:oob")
	q.Set("response_type", "code")
	q.Set("scope", "public")

	authURL.RawQuery = q.Encode()

	if err := openbrowser(authURL.String()); err != nil {
		return "", err
	}

	fmt.Println("Enter authorization code:")

	var code string

	if _, err := fmt.Scanln(&code); err != nil {
		return "", err
	}

	authorizedCode, err := a.authorizeCode(code)
	if err != nil {
		return "", err
	}

	token := authorizedCode.AccessToken

	if len(token) < 1 {
		return "", fmt.Errorf("broken token response from unsplash")
	}

	if err := a.tokenStore.Set([]byte(token)); err != nil {
		return "", err
	}

	return token, nil
}

type authorizeCodeResponse struct {
	AccessToken string `json:"access_token"`
}

func (a *UnsplashAuthorizer) authorizeCode(code string) (*authorizeCodeResponse, error) {
	requestURL := url.URL{
		Scheme: "https",
		Host:   "unsplash.com",
		Path:   "/oauth/token",
	}
	q := requestURL.Query()
	q.Set("client_id", a.clientID)
	q.Set("client_secret", a.clientSecret)
	q.Set("code", code)
	q.Set("redirect_uri", "urn:ietf:wg:oauth:2.0:oob")
	q.Set("grant_type", "authorization_code")

	requestURL.RawQuery = q.Encode()

	resp, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var data interface{}

		json.NewDecoder(resp.Body).Decode(&data)

		return nil, fmt.Errorf("code authorization return non ok code: %d. %v", resp.StatusCode, data)
	}

	var result authorizeCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *UnsplashAuthorizer) registerUser(accessToken string) (string, error) {
	reqURL := url.URL{
		Scheme: "https",
		Host:   "api.unsplash.com",
		Path:   "/clients",
	}

	data := url.Values{
		"name":        {"Wallpaperize"},
		"description": {"Application for setting wallpapers from various sources"},
	}

	res, err := http.PostForm(reqURL.String(), data)
	if err != nil {
		return "", err
	}

	type response struct {
		ClientID string `json:"client_id"`
	}

	var result response

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.ClientID, nil
}

func openbrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}
