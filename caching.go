package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"time"

	"github.com/goncharovnikita/wallpaperize/api"
)

var (
	cacheLimit = 10
)

func cache() {
	var (
		err   error
		names []string
	)

	if names, err = conf.parseConfig(); err != nil {
		log.Fatal(err)
	}

	if len(names) < cacheLimit {
		var u api.UnsplashAPI
		var name string
		for i := 0; i < (cacheLimit - len(names)); i++ {
			if name, err = saveToCache(u); err != nil {
				log.Fatal(err)
			}
			if err = conf.add(name); err != nil {
				log.Fatal(err)
			}
		}
	}
}

// saveToCache performs getting image and
// saving to cache
func saveToCache(imageGetter RandomImageReader) (name string, err error) {
	var (
		result io.ReadCloser
		img    image.Image
		file   *os.File
		format string
	)

	if result, err = imageGetter.GetRandomImageReader(); err != nil {
		return
	}

	defer result.Close()

	if img, format, err = image.Decode(result); err != nil {
		return
	}

	name = string(time.Now().Format(time.UnixDate)) + "." + format

	defer file.Close()

	if file, err = os.OpenFile(absCacheDirname+"/"+name, os.O_CREATE|os.O_RDWR, fileperm); err != nil {
		return
	}

	switch format {
	case "jpeg":
		if err = jpeg.Encode(file, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			return
		}
		break
	case "png":
		if err = png.Encode(file, img); err != nil {
			return
		}
		break
	default:
		log.Printf("Unknown format %s\n", format)
	}

	return
}
