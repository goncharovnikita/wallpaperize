package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	picsumAPIprefix      = "https://picsum.photos/"
	picsumRandomPhotoURL = "1920/1080/?random"
)

// PicsumAPI implementation
type PicsumAPI struct{}

// GetRandomImage implementation
func (u PicsumAPI) GetRandomImage() (result []byte) {
	var (
		response *http.Response
		body     []byte
		client   http.Client
		request  *http.Request
		URL      string
		err      error
	)

	URL = picsumAPIprefix + picsumRandomPhotoURL

	if request, err = http.NewRequest(http.MethodGet, URL, nil); err != nil {
		log.Fatal(err)
	}

	if response, err = client.Do(request); err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatal(err)
	}

	return body
}
