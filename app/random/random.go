package random

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/hash"
	"github.com/goncharovnikita/wallpaperize/back/models"
)

type randomImageGetter interface {
	GetRandomImages(limit int) ([]*models.UnsplashImage, error)
}

// Random image getter type
type Random struct {
	rg   randomImageGetter
	path string
}

// NewRandomImageGetter create new random image getter
func NewRandomImageGetter(rg randomImageGetter, path string) *Random {
	return &Random{
		rg:   rg,
		path: path,
	}
}

// GetImage implementation
func (r *Random) GetImage() (string, error) {
	img, err := r.rg.GetRandomImages(1)
	if err != nil {
		return "", err
	}

	if len(img) < 1 {
		return "", fmt.Errorf("empty images response")
	}

	resp, err := http.Get(img[0].URLs.Full)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	hsh, err := hash.Hash256(b)
	if err != nil {
		return "", err
	}

	fname := r.path + "/random" + hsh + ".jpg"

	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	return fname, nil
}
