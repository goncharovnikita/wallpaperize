package updater

import (
	"log"
	"sync"
	"time"

	"github.com/goncharovnikita/wallpaperize/app/api"
	"github.com/goncharovnikita/wallpaperize/back/models"
)

type imagesSetter interface {
	SetImages([]*models.UnsplashImage) error
}

type Unsplash struct {
	api *api.UnsplashAPI
	imagesSetter
	logger     *log.Logger
	shouldStop bool
	mux        *sync.Mutex
}

func NewUnsplash(
	accessToken string,
	imagesSetter imagesSetter,
	logger *log.Logger,
) *Unsplash {
	return &Unsplash{
		api: api.NewUnsplashAPI(staticTokenGetter{
			token: accessToken,
		}),
		imagesSetter: imagesSetter,
		logger:       logger,
		shouldStop:   false,
		mux:          &sync.Mutex{},
	}
}

func (u *Unsplash) Run() error {
	for {
		images := make([]*models.UnsplashImage, 0)
		ctr := 2

		for {
			u.mux.Lock()

			if u.shouldStop {
				u.mux.Unlock()

				return nil
			}

			u.mux.Unlock()

			img, err := u.api.GetRandomImage()
			if err != nil {
				u.logger.Printf("getting random image from unsplash, %v\n", err)

				break
			}

			images = append(images, models.MakeUnsplashImageFromAPI(&img.Data))

			u.logger.Printf("saved: %d, to proceed: %d\n", len(images), img.RateLimitRemaining)

			if img.RateLimitRemaining < 1 || ctr <= 0 {
				break
			}

			ctr--
		}

		if len(images) > 0 {
			if err := u.imagesSetter.SetImages(images); err != nil {
				u.logger.Printf("error setting images to db, %v\n", err)

				break
			}

			u.logger.Printf("saved %d images\n", len(images))
		}

		<-time.NewTimer(1 * time.Hour).C
	}

	return nil
}

func (u *Unsplash) Stop() error {
	u.mux.Lock()

	u.shouldStop = true

	u.mux.Unlock()

	return nil
}

type staticTokenGetter struct {
	token string
}

func (s staticTokenGetter) GetToken() (string, error) {
	return s.token, nil
}
