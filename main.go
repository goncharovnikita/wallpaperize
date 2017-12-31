package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
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
		tempFilename = "/bing_daily_image.png"
		cacheDirname = "/.wallpaperize_cache"
		usr          *user.User
		absPath      string
		err          error
		file         *os.File
	)

	dailyImage = bingAPI.GetDailyImage()

	if usr, err = user.Current(); err != nil {
		log.Fatal(err)
	}

	if err = os.Mkdir(usr.HomeDir+cacheDirname, 0777); err != nil {
		var ok bool
		if _, ok = err.(*os.PathError); !ok {
			log.Fatal(err)
		}
	}

	if file, err = os.OpenFile(usr.HomeDir+cacheDirname+tempFilename, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err = file.Write(dailyImage); err != nil {
		log.Fatal(err)
	}

	if absPath, err = filepath.Abs(usr.HomeDir + cacheDirname + tempFilename); err != nil {
		log.Fatal(err)
	}

	if err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", absPath).Run(); err != nil {
		log.Fatal(err)
	}
}
