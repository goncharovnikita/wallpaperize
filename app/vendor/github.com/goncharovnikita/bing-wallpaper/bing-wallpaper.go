package bing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var apiURL = "http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
var imageURLHost = "http://www.bing.com/"

type Client struct {
	c *http.Client
}

func NewClient(c *http.Client) *Client {
	return &Client{
		c: c,
	}
}

type Image struct {
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
	Images []*Image `json:"images"`
}

// GetDailyImage get daily image based on contry code
// For example: en-EN
func (b *Client) GetDailyImage(countryCode string) (*Image, error) {
	res, err := b.c.Get(fmt.Sprintf("%s&mkt=%s", apiURL, countryCode))
	if err != nil {
		return nil, err
	}

	var data getImageInfoResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Images) < 1 {
		return nil, fmt.Errorf("no images in response")
	}

	return data.Images[0], nil
}

// GetDailyImageRaw get daily image raw data based on contry code
// For example: en-EN
func (b *Client) GetDailyImageRaw(countryCode string) ([]byte, error) {
	img, err := b.GetDailyImage(countryCode)
	if err != nil {
		return nil, err
	}

	url := img.URL
	if len(url) < 1 {
		return nil, fmt.Errorf("image url length is empty")
	}

	res, err := b.c.Get(imageURLHost + url)
	if err != nil {
		return nil, err
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
