package updater

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/goncharovnikita/unsplash"
	"github.com/goncharovnikita/wallpaperize/back/mapper"
	"github.com/goncharovnikita/wallpaperize/back/models"
)

const limitPerLoop = 50
const maxImagesCount = 3000
const removeImagesCount = 1000

type imagesSetter interface {
	SetImages([]*models.UnsplashImage) error
}

type repoCleaner interface {
	ImagesCount() (int, error)
	RemoveFirstImages(count int) error
}

type Unsplash struct {
	api          *unsplash.API
	imagesSetter imagesSetter
	repoCleaner  repoCleaner
	logger       *log.Logger
	shouldStop   bool
	mux          *sync.Mutex
}

func NewUnsplash(
	accessToken string,
	imagesSetter imagesSetter,
	repoCleaner repoCleaner,
	logger *log.Logger,
) *Unsplash {
	return &Unsplash{
		api: unsplash.NewUnsplashAPI(staticTokenGetter{
			token: accessToken,
		}, http.DefaultClient),
		imagesSetter: imagesSetter,
		repoCleaner:  repoCleaner,
		logger:       logger,
		shouldStop:   false,
		mux:          &sync.Mutex{},
	}
}

func (u *Unsplash) Run() error {
	for {
		u.mux.Lock()

		if u.shouldStop {
			u.mux.Unlock()

			return nil
		}

		u.mux.Unlock()

		imagesCount, err := u.repoCleaner.ImagesCount()
		if err != nil {
			u.logger.Printf("error getting items count: %v", err)

			continue
		}

		if imagesCount >= maxImagesCount {
			u.logger.Printf("removing %d images", removeImagesCount)

			if err := u.repoCleaner.RemoveFirstImages(removeImagesCount); err != nil {
				u.logger.Printf("error removing repo items: %v", err)

				continue
			}

			u.logger.Printf("successfully removed %d images", removeImagesCount)
		}

		images := make([]*models.UnsplashImage, 0)
		ctr := limitPerLoop
		remaining := 50

		for {
			u.mux.Lock()

			if u.shouldStop {
				u.mux.Unlock()

				return nil
			}

			u.mux.Unlock()

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			img, err := u.api.GetRandomImage(ctx, "landscape", "1920", "1080")
			if err != nil {
				u.logger.Printf("getting random image from unsplash, %v\n", err)

				break
			}

			if img.RateLimitRemaining == remaining {
				continue
			}

			remaining = img.RateLimitRemaining

			images = append(images, mapper.MakeUnsplashImageFromAPI(&img.Data))

			u.logger.Printf("saved: %d, to proceed: %d\n", len(images), img.RateLimitRemaining)

			if img.RateLimitRemaining < 1 || ctr <= 1 {
				break
			}

			ctr--

			// Unsplash caches sequential requests
			time.Sleep(1 * time.Second)
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
