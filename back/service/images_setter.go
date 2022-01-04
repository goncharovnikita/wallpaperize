package service

import "github.com/goncharovnikita/wallpaperize/back/models"

type imagesSetterRepo interface {
	SetImages([]*models.DBImage) error
}

type ImagesSetter struct {
	repo imagesSetterRepo
}

func (s *ImagesSetter) SetImages(images []*models.ResponseImage) error {
	result := make([]*models.DBImage, 0, len(images))

	for _, image := range images {
		result = append(result, models.MakeDBImage(image))
	}

	if err := s.repo.SetImages(result); err != nil {
		return err
	}

	return nil
}
