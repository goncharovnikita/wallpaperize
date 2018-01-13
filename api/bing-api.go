package api

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var apiURL = "http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US"
var apiPrefix = "http://www.bing.com/"

// BingAPI DailyImageGetter implementation
type BingAPI struct{}

type imageInfo struct {
	Startdate     string `json:"startdate"`
	Fullstartdate string `json:"fullstartdate"`
	Enddate       string `json:"enddate"`
	URL           string `json:"url"`
	Urlbase       string `json:"urlbase"`
	Copyright     string `json:"copyright"`
	Copyrightlink string `json:"copyrightlink"`
	Quiz          string `json:"quiz"`
	Wp            bool   `json:"wp"`
	Hsh           string `json:"hsh"`
}

type getImageInfoResponse struct {
	Images []imageInfo `json:"images"`
}

// GetDailyImage implementation
func (b BingAPI) GetDailyImage() (result []byte) {
	var (
		response *http.Response
		data     getImageInfoResponse
		body     []byte
		err      error
		url      string
		raw      interface{}
	)

	if response, err = http.Get(apiURL); err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &raw); err != nil {
		log.Fatal(err)
	}

	if len(data.Images) < 1 {
		log.Printf("%+v\n", raw)
		log.Printf("%+v\n", data)
		log.Fatal("images size less than 1")
	}

	if url = data.Images[0].URL; len(url) < 1 {
		log.Fatal("url len less than 1")
	}

	if response, err = http.Get(apiPrefix + url); err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatal(err)
	}

	if _, err = base64.StdEncoding.Decode(result, body); err != nil {
		log.Fatal(err)
	}

	return
}

// GetImageReader implementation
func (b BingAPI) GetImageReader() (result io.ReadCloser, err error) {
	var (
		response *http.Response
		data     getImageInfoResponse
		url      string
		body     []byte
	)

	if response, err = http.Get(apiURL); err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	if len(data.Images) < 1 {
		log.Printf("%+v\n", data)
		log.Fatal("images size less than 1")
	}

	if url = data.Images[0].URL; len(url) < 1 {
		log.Fatal("url len less than 1")
	}

	if response, err = http.Get(apiPrefix + url); err != nil {
		log.Fatal(err)
	}

	result = response.Body
	return
}
