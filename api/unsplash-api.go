package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	unsplashAPIprefix      = "https://api.unsplash.com/"
	unsplashRandomPhotoURL = "photos/random"
	unsplashAppID          = "288c71e3029fe7ff9572e518dfce06b383676fb0a7c1d8bc10cc3e06af252ed5"
	unsplashAppIDBackup    = "af7ad21e753b2d1bac5e0d63d1ce9e1abefe5facb41c6252f1c741a45b64a658"
)

// UnsplashAPI implementation
type UnsplashAPI struct{}

type unsplashRandomImageURLs struct {
	RAW string `json:"raw"`
}

type unsplashRandomImageResponse struct {
	ID          string                  `json:"id"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
	Width       int                     `json:"width"`
	Height      int                     `json:"height"`
	Description string                  `json:"description"`
	URLs        unsplashRandomImageURLs `json:"urls"`
}

// GetRandomImage implementation
func (u UnsplashAPI) GetRandomImage() (result []byte, err error) {
	var (
		response *http.Response
		data     unsplashRandomImageResponse
		body     []byte
		client   http.Client
		request  *http.Request
		URL      string
	)

	URL = unsplashAPIprefix + unsplashRandomPhotoURL + "?orientation=landscape&w=1920&h=1080"

	if request, err = http.NewRequest(http.MethodGet, URL, nil); err != nil {
		return
	}

	request.Header.Set("Authorization", "Client-ID "+unsplashAppID)

	if response, err = client.Do(request); err != nil {
		request.Header.Del("Authorization")
		request.Header.Set("Authorization", "Client-ID "+unsplashAppIDBackup)
		if response, err = client.Do(request); err != nil {
			log.Print(err)
			return
		}
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &data); err != nil {
		log.Print(err)
		return
	}

	if len(data.URLs.RAW) < 1 {
		log.Printf("%+v\n", data)
		log.Fatal("raw url len less than 1")
	}

	if response, err = http.Get(data.URLs.RAW); err != nil {
		log.Print(err)
		return
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Print(err)
		return
	}

	return body, nil
}

// GetImageReader returns an Reader with
// image entity
func (u UnsplashAPI) GetImageReader() (result io.ReadCloser, err error) {
	var (
		response *http.Response
		data     unsplashRandomImageResponse
		body     []byte
		client   http.Client
		request  *http.Request
		URL      string
	)

	URL = unsplashAPIprefix + unsplashRandomPhotoURL + "?orientation=landscape&w=1920&h=1080"

	if request, err = http.NewRequest(http.MethodGet, URL, nil); err != nil {
		log.Print(err)
		return
	}

	request.Header.Set("Authorization", "Client-ID "+unsplashAppID)

	if response, err = client.Do(request); err != nil {
		log.Print(err)
		request.Header.Del("Authorization")
		request.Header.Set("Authorization", "Client-ID "+unsplashAppIDBackup)
		if response, err = client.Do(request); err != nil {
			log.Print(err)
			return
		}
	}

	if response.StatusCode != http.StatusOK {
		request.Header.Del("Authorization")
		request.Header.Set("Authorization", "Client-ID "+unsplashAppIDBackup)
		if response, err = client.Do(request); err != nil {
			log.Print(err)
			return
		}
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Print(err)
		return
	}

	if err = json.Unmarshal(body, &data); err != nil {
		log.Print(string(body), err)
		return
	}

	if len(data.URLs.RAW) < 1 {
		log.Printf("%+v\n", data)
		log.Fatal("raw url len less than 1")
	}

	if response, err = http.Get(data.URLs.RAW); err != nil {
		log.Print(err)
		return
	}

	result = response.Body
	return
}
