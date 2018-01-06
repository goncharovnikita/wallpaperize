package main

import (
	"io/ioutil"
	"net/http"
)

var (
	picsumAPIprefix      = "https://picsum.photos/"
	picsumRandomPhotoURL = "1920/1080/?random"
)

// PicsumAPI implementation
type PicsumAPI struct{}

// GetRandomImage implementation
func (u PicsumAPI) GetRandomImage() (result []byte, err error) {
	var (
		response *http.Response
		body     []byte
		client   http.Client
		request  *http.Request
		URL      string
	)

	URL = picsumAPIprefix + picsumRandomPhotoURL

	if request, err = http.NewRequest(http.MethodGet, URL, nil); err != nil {
		return
	}

	if response, err = client.Do(request); err != nil {
		return
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}

	return body, nil
}
