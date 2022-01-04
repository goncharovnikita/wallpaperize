package random

import (
	"bytes"
	"io"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/hash"
)

type randomImageGetter interface {
	GetImage() ([]byte, error)
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
	img, err := r.rg.GetImage()
	if err != nil {
		return "", err
	}

	hsh, err := hash.Hash256(img)
	if err != nil {
		return "", err
	}

	fname := r.path + "/random" + hsh + ".jpg"

	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(img))
	if err != nil {
		return "", err
	}

	return fname, nil
}
