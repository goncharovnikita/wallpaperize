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
	cacheLimit     = 10
	staticFilename = "/bing_daily_image.jpeg"
)

// cache type
type cache struct {
	absPath string
}

func (c cache) cache() {
	var (
		err   error
		struc configStructure
	)

	if struc, err = conf.parseConfig(); err != nil {
		log.Fatal(err)
	}

	if len(struc.RandomPhotos) < cacheLimit {
		var u api.UnsplashAPI
		var name string
		for i := 0; i < (cacheLimit - len(struc.RandomPhotos)); i++ {
			if name, err = c.saveToCache(u, true); err != nil {
				log.Fatal(err)
			}
			if err = conf.addRandomPhoto(name); err != nil {
				log.Fatal(err)
			}
		}
	}

	if _, ex := struc.DailyImages["bing"]; !ex {
		var u api.BingAPI
		if _, err = c.saveToCache(u, true); err != nil {
			log.Fatal(err)
		}
	} else {
		now := time.Now()
		t, e := time.Parse("2006-01-02", struc.DailyImages["bing"].Date)

		if e != nil {
			log.Fatal(e)
		}

		if !compareDates(&now, &t) {
			var u api.BingAPI
			if _, err = c.saveToCache(u, true); err != nil {
				log.Fatal(err)
			}
		}
	}
}

// saveToCache performs getting image and
// saving to cache
func (c cache) saveToCache(imageGetter ImageReaderGetter, random bool) (name string, err error) {
	var (
		result  io.ReadCloser
		img     image.Image
		file    *os.File
		format  string
		dirname string
	)

	if result, err = imageGetter.GetImageReader(); err != nil {
		log.Print(err)
		return
	}

	defer result.Close()

	if img, format, err = image.Decode(result); err != nil {
		log.Print(err)
		return
	}

	if random {
		name = string(time.Now().Format(time.UnixDate)) + "." + format
		dirname = absRandomDirname + "/" + name
	} else {
		name = staticFilename
		dirname = absCacheDirname + "/" + name
	}

	defer file.Close()

	if file, err = os.OpenFile(dirname, os.O_CREATE|os.O_RDWR, fileperm); err != nil {
		log.Print(err)
		return
	}

	switch format {
	case "jpeg":
		if err = jpeg.Encode(file, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			log.Print(err)
			return
		}
		break
	case "png":
		if err = png.Encode(file, img); err != nil {
			log.Print(err)
			return
		}
		break
	default:
		log.Printf("Unknown format %s\n", format)
	}

	return
}

// retrieves data from cache if
// exists
func (c cache) retrieve(random bool) (result string, err error) {
	if !random {
		return c.retrieveStatic()
	}
	return c.retrieveNonStatic()
}

// retrieve photo not random photo
// from the cache
func (c cache) retrieveStatic() (result string, err error) {
	if _, err = os.OpenFile(absCacheDirname+staticFilename, os.O_RDONLY, fileperm); err != nil {
		return
	}
	result = staticFilename
	return
}

// retrieve cached random photo
func (c cache) retrieveNonStatic() (string, error) {
	return conf.switchNext()
}
