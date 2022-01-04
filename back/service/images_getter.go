package service

import "github.com/goncharovnikita/wallpaperize/back/models"

type imagesGetterRepo interface {
	GetImages(limit int) ([]*models.DBImage, error)
}

type ImagesGetter struct {
	repo imagesGetterRepo
}

func NewImagesGetter(repo imagesGetterRepo) *ImagesGetter {
	return &ImagesGetter{
		repo: repo,
	}
}

func (s *ImagesGetter) GetImages(limit int) ([]*models.ResponseImage, error) {
	images, err := s.repo.GetImages(limit)
	if err != nil {
		return nil, err
	}

	result := make([]*models.ResponseImage, 0, len(images))

	for _, img := range images {
		result = append(result, models.MakeResponseImage(img))
	}

	return result, nil
}
