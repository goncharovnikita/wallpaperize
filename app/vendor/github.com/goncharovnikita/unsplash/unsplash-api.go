package unsplash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	apiPrefix       = "https://api.unsplash.com/"
	randomPhotoPath = "photos/random"
)

type idGetter interface {
	GetToken() (string, error)
}

// API implementation
type API struct {
	idGetter   idGetter
	httpClient *http.Client
}

func NewUnsplashAPI(
	idGetter idGetter,
	httpClient *http.Client,
) *API {
	return &API{
		idGetter:   idGetter,
		httpClient: httpClient,
	}
}

type RandomImageURLs struct {
	RAW     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type ImageLinks struct {
	Self             string `json:"self"`
	Html             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}

type Image struct {
	ID          string          `json:"id"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Width       int             `json:"width"`
	Height      int             `json:"height"`
	Description string          `json:"description"`
	URLs        RandomImageURLs `json:"urls"`
	Links       ImageLinks      `json:"links"`
}

type ImageResponse struct {
	Data               Image
	RateLimitRemaining int
}

func (u *API) GetRandomImage(
	orientation string,
	width string,
	height string,
) (*ImageResponse, error) {
	token, err := u.idGetter.GetToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(
		"%s%s?orientation=%s&w=%s&h=%s",
		apiPrefix,
		randomPhotoPath,
		orientation,
		width,
		height,
	)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Client-ID "+token)

	response, err := u.httpClient.Do(request)
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

	var data Image

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &ImageResponse{
		Data:               data,
		RateLimitRemaining: rateLimitRemaining,
	}, nil
}
