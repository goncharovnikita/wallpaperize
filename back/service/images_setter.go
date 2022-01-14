package service

import (
	"github.com/goncharovnikita/wallpaperize/back/internal/models"
	"github.com/goncharovnikita/wallpaperize/back/mapper"
	pubmodels "github.com/goncharovnikita/wallpaperize/back/models"
)

type imagesSetterRepo interface {
	SetImages([]*models.DBImage) error
}

type ImagesSetter struct {
	repo imagesSetterRepo
}

func NewImagesSetter(repo imagesSetterRepo) *ImagesSetter {
	return &ImagesSetter{
		repo: repo,
	}
}

func (s *ImagesSetter) SetImages(images []*pubmodels.UnsplashImage) error {
	result := make([]*models.DBImage, 0, len(images))

	for _, image := range images {
		img, err := mapper.MakeDBImage(image)
		if err != nil {
			return err
		}

		result = append(result, img)
	}

	if err := s.repo.SetImages(result); err != nil {
		return err
	}

	return nil
}
