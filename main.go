package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var dailyImage []byte
var bingAPI BingAPI

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	setPhotoOfTheDay()
}

func setPhotoOfTheDay() {
	var (
		tempFilename = "bing_daily_image.png"
		absPath      string
		err          error
		file         *os.File
	)

	dailyImage = bingAPI.GetDailyImage()

	if file, err = os.Create(tempFilename); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err = file.Write(dailyImage); err != nil {
		log.Fatal(err)
	}

	if absPath, err = filepath.Abs("./" + tempFilename); err != nil {
		log.Fatal(err)
	}

	if err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", absPath).Run(); err != nil {
		log.Fatal(err)
	}
}
