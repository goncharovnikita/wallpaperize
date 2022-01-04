package models

import (
	"encoding/json"

	"github.com/goncharovnikita/wallpaperize/app/api"
)

type DBImage struct {
	Data []byte
}

type UnsplashImageURL struct {
	RAW     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type UnsplashImageLinks struct {
	Self             string `json:"self"`
	Html             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}

type UnsplashImage struct {
	ID          string             `json:"id"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
	Width       int                `json:"width"`
	Height      int                `json:"height"`
	Description string             `json:"description"`
	URLs        UnsplashImageURL   `json:"urls"`
	Links       UnsplashImageLinks `json:"links"`
}

func MakeUnsplashImageFromAPI(image *api.UnsplashRandomImage) *UnsplashImage {
	return &UnsplashImage{
		ID:          image.ID,
		CreatedAt:   image.CreatedAt,
		UpdatedAt:   image.UpdatedAt,
		Width:       image.Width,
		Height:      image.Height,
		Description: image.Description,
		URLs: UnsplashImageURL{
			RAW:     image.URLs.RAW,
			Full:    image.URLs.Full,
			Regular: image.URLs.Regular,
			Small:   image.URLs.Small,
			Thumb:   image.URLs.Thumb,
		},
		Links: UnsplashImageLinks{
			Self:             image.Links.Self,
			Html:             image.Links.Html,
			Download:         image.Links.Download,
			DownloadLocation: image.Links.DownloadLocation,
		},
	}
}

type ImagesResponse struct {
	Data []*UnsplashImage `json:"data"`
}

func MakeUnsplashImage(image *DBImage) (*UnsplashImage, error) {
	var data UnsplashImage
	if err := json.Unmarshal(image.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func MakeDBImage(image *UnsplashImage) (*DBImage, error) {
	data, err := json.Marshal(image)
	if err != nil {
		return nil, err
	}

	return &DBImage{
		Data: data,
	}, nil
}

type ResponseError struct {
	Error string `json:"error"`
}
