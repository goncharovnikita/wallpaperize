package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goncharovnikita/wallpaperize/back/models"
)

type HTTP struct {
	c       http.Client
	baseURL string
}

func NewHTTP(baseURL string) *HTTP {
	return &HTTP{
		c:       *http.DefaultClient,
		baseURL: baseURL,
	}
}

func (c *HTTP) GetRandomImages(limit int) ([]*models.UnsplashImage, error) {
	url := fmt.Sprintf("%s/get/random?limit=%d", c.baseURL, limit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var result []*models.UnsplashImage

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
