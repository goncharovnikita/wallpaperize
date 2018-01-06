package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var unsplashRandomFlag = flag.Bool("u", false, "set random picture from unsplash as wallpaper")
var picsumRandomFlag = flag.Bool("p", false, "set random picture from lorem picsum as wallpaper")
var cacheDirname = "/.wallpaperize_cache"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()
	switch true {
	case !isGNOMECompatible():
		fmt.Printf("cannot perform wallpaperize - such system is not compatible yet\n")
		break
	case *unsplashRandomFlag != false:
		var unsplashAPI UnsplashAPI
		setRandomPhoto(unsplashAPI)
		break
	case *picsumRandomFlag != false:
		var picsumAPI PicsumAPI
		setRandomPhoto(picsumAPI)
		break
	default:
		setPhotoOfTheDay()
		break
	}
}

// RandomImageGetter interface provide get random inage method
type RandomImageGetter interface {
	GetRandomImage() ([]byte, error)
}

func setPhotoOfTheDay() {
	var (
		tempFilename = "/bing_daily_image.png"
		absPath      string
		err          error
		file         *os.File
		dailyImage   []byte
		bingAPI      BingAPI
	)

	createCacheFolder()

	dailyImage = bingAPI.GetDailyImage()

	if file, err = os.OpenFile(getAbsCacheDirname()+tempFilename, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err = file.Write(dailyImage); err != nil {
		log.Fatal(err)
	}

	if absPath, err = filepath.Abs(getAbsCacheDirname() + tempFilename); err != nil {
		log.Fatal(err)
	}

	if err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", absPath).Run(); err != nil {
		log.Fatal(err)
	}
}

func setRandomPhoto(imageGetter RandomImageGetter) {
	var (
		tempFilename = "/random_image.png"
		absPath      string
		err          error
		file         *os.File
		randomImage  []byte
	)

	createCacheFolder()

	if randomImage, err = imageGetter.GetRandomImage(); err != nil {
		log.Fatal(err)
	}

	if file, err = os.OpenFile(getAbsCacheDirname()+tempFilename, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err = file.Write(randomImage); err != nil {
		log.Fatal(err)
	}

	if absPath, err = filepath.Abs(getAbsCacheDirname() + tempFilename); err != nil {
		log.Fatal(err)
	}

	if err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", absPath).Run(); err != nil {
		log.Fatal(err)
	}
}
