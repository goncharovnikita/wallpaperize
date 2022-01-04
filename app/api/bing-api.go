package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var apiURL = "http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
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
func (b BingAPI) GetDailyImage() ([]byte, error) {
	client := http.Client{}

	req1, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	qu := req1.URL.Query()
	qu.Add("mkt", "ru-RU")
	req1.URL.RawQuery = qu.Encode()

	response1, err := client.Do(req1)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer response1.Body.Close()

	body1, err := io.ReadAll(response1.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var data getImageInfoResponse

	if err = json.Unmarshal(body1, &data); err != nil {
		log.Println(err)
		return nil, err
	}

	if len(data.Images) < 1 {
		log.Printf("%+v\n", data)
		log.Fatal("images size less than 1")
	}

	url := data.Images[0].URL
	if len(url) < 1 {
		return nil, fmt.Errorf("url len less than 1")
	}

	response2, err := http.Get(apiPrefix + url)
	if err != nil {
		return nil, err
	}

	defer response2.Body.Close()

	body2, err := io.ReadAll(response2.Body)
	if err != nil {
		return nil, err
	}

	return body2, nil
}
