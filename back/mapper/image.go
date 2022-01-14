package mapper

import (
	"encoding/json"

	"github.com/goncharovnikita/unsplash"
	"github.com/goncharovnikita/wallpaperize/back/internal/models"
	pubmodels "github.com/goncharovnikita/wallpaperize/back/models"
)

func MakeUnsplashImageFromAPI(image *unsplash.Image) *pubmodels.UnsplashImage {
	return &pubmodels.UnsplashImage{
		ID:          image.ID,
		CreatedAt:   image.CreatedAt,
		UpdatedAt:   image.UpdatedAt,
		Width:       image.Width,
		Height:      image.Height,
		Description: image.Description,
		URLs: pubmodels.UnsplashImageURL{
			RAW:     image.URLs.RAW,
			Full:    image.URLs.Full,
			Regular: image.URLs.Regular,
			Small:   image.URLs.Small,
			Thumb:   image.URLs.Thumb,
		},
		Links: pubmodels.UnsplashImageLinks{
			Self:             image.Links.Self,
			Html:             image.Links.Html,
			Download:         image.Links.Download,
			DownloadLocation: image.Links.DownloadLocation,
		},
	}
}

func MakeUnsplashImage(image *models.DBImage) (*pubmodels.UnsplashImage, error) {
	var data pubmodels.UnsplashImage
	if err := json.Unmarshal(image.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func MakeDBImage(image *pubmodels.UnsplashImage) (*models.DBImage, error) {
	data, err := json.Marshal(image)
	if err != nil {
		return nil, err
	}

	return &models.DBImage{
		Data: data,
	}, nil
}
