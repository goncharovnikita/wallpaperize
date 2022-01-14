package service

import (
	"github.com/goncharovnikita/wallpaperize/back/internal/models"
	"github.com/goncharovnikita/wallpaperize/back/mapper"
	pubmodels "github.com/goncharovnikita/wallpaperize/back/models"
)

type imagesGetterRepo interface {
	GetRandomImages(limit int) ([]*models.DBImage, error)
}

type ImagesGetter struct {
	repo imagesGetterRepo
}

func NewImagesGetter(repo imagesGetterRepo) *ImagesGetter {
	return &ImagesGetter{
		repo: repo,
	}
}

func (s *ImagesGetter) GetImages(limit int) ([]*pubmodels.UnsplashImage, error) {
	images, err := s.repo.GetRandomImages(limit)
	if err != nil {
		return nil, err
	}

	result := make([]*pubmodels.UnsplashImage, 0, len(images))

	for _, image := range images {
		img, err := mapper.MakeUnsplashImage(image)
		if err != nil {
			return nil, err
		}

		result = append(result, img)
	}

	return result, nil
}
