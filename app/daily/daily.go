// Package daily encapsulates work with daily images
package daily

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

// Daily fetch daily images
type Daily struct {
	dg   DailyGetter
	path string
}

// DailyGetter daily getter interface
type DailyGetter interface {
	GetDailyImageRaw(ctx context.Context, countryCode string) ([]byte, error)
}

// NewDailyGetter creates new daily getter instance
func NewDailyGetter(dg DailyGetter, cachePath string) *Daily {
	return &Daily{
		dg:   dg,
		path: cachePath,
	}
}

// GetImage ImageGetter implementation
func (d *Daily) GetImage() (string, error) {
	name := d.path + "/" + getToday() + ".jpg"

	info, err := os.Stat(name)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			img, err := d.dg.GetDailyImageRaw(ctx, "en-EN")
			if err != nil {
				return "", err
			}

			file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				return "", err
			}

			defer file.Close()

			_, err = file.Write(img)
			if err != nil {
				return "", err
			}

			return name, nil
		}

		return "", err
	}

	if info.IsDir() {
		return "", fmt.Errorf("path is directory, please delete and retry again - %s", name)
	}

	return name, nil
}

func getToday() string {
	year, month, day := time.Now().Date()

	return fmt.Sprintf("%d.%s.%d", day, month, year)
}
