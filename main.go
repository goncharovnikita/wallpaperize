package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var randomFlag = flag.Bool("r", false, "set random picture as wallpaper")
var cacheDirname = "/.wallpaperize_cache"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	if !isGNOMECompatible() {
		fmt.Printf("cannot perform wallpaperize - such system is not compatible yet\n")
		return
	}
	flag.Parse()
	if *randomFlag != false {
		setRandomPhoto()
		return
	}
	setPhotoOfTheDay()
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

func setRandomPhoto() {
	var (
		tempFilename = "/random_image.png"
		absPath      string
		err          error
		file         *os.File
		randomImage  []byte
		unsplashAPI  UnsplashAPI
	)

	createCacheFolder()

	randomImage = unsplashAPI.GetRandomImage()

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
