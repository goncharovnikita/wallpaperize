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

func (s *ImagesGetter) GetImages(limit int) ([]*models.UnsplashImage, error) {
	images, err := s.repo.GetImages(limit)
	if err != nil {
		return nil, err
	}

	result := make([]*models.UnsplashImage, 0, len(images))

	for _, image := range images {
		img, err := models.MakeUnsplashImage(image)
		if err != nil {
			return nil, err
		}

		result = append(result, img)
	}

	return result, nil
}
