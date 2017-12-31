package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	unsplashAPIprefix      = "https://api.unsplash.com/"
	unsplashRandomPhotoURL = "photos/random"
	unsplashAppID          = "288c71e3029fe7ff9572e518dfce06b383676fb0a7c1d8bc10cc3e06af252ed5"
)

// RandomImageGetter interface provide get random inage method
type RandomImageGetter interface {
	GetRandomImage() ([]byte, error)
}

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
func (u UnsplashAPI) GetRandomImage() (result []byte) {
	var (
		response *http.Response
		data     unsplashRandomImageResponse
		body     []byte
		client   http.Client
		request  *http.Request
		URL      string
		err      error
	)

	URL = unsplashAPIprefix + unsplashRandomPhotoURL + "?orientation=landscape&w=1920&h=1080"

	if request, err = http.NewRequest(http.MethodGet, URL, nil); err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Authorization", "Client-ID "+unsplashAppID)

	if response, err = client.Do(request); err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	if len(data.URLs.RAW) < 1 {
		log.Printf("%+v\n", data)
		log.Fatal("raw url len less than 1")
	}

	if response, err = http.Get(data.URLs.RAW); err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatal(err)
	}

	return body
}
