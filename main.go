package main

import (
	"encoding/base64"
	"flag"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var unsplashRandomFlag = flag.Bool("u", false, "set random picture from unsplash as wallpaper")
var picsumRandomFlag = flag.Bool("p", false, "set random picture from lorem picsum as wallpaper")
var daemonizeFlag = flag.Bool("d", false, "run as daemon")
var cacheDirname = "/.wallpaperize_cache"
var absCacheDirname string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	createCacheFolder()
}

func main() {
	cache()
	// flag.Parse()
	// switch true {
	// case !isGNOMECompatible():
	// 	fmt.Printf("cannot perform wallpaperize - such system is not compatible yet\n")
	// 	break
	// case *unsplashRandomFlag != false:
	// 	var unsplashAPI api.UnsplashAPI
	// 	setRandomPhoto(unsplashAPI)
	// 	break
	// case *picsumRandomFlag != false:
	// 	var picsumAPI PicsumAPI
	// 	setRandomPhoto(picsumAPI)
	// 	break
	// case *daemonizeFlag != false:
	// 	runAsDaemon()
	// 	break
	// default:
	// 	setPhotoOfTheDay()
	// 	break
	// }
}

func setPhotoOfTheDay() {
	var (
		tempFilename     = "/bing_daily_image"
		absPath          string
		err              error
		file             *os.File
		dailyImageReader io.Reader
		bingAPI          BingAPI
		img              image.Image
		format           string
	)

	if dailyImageReader, err = bingAPI.GetDailyImageReader(); err != nil {
		log.Fatal(err)
	}

	if img, format, err = image.Decode(dailyImageReader); err != nil {
		log.Fatal(err)
	}

	tempFilename = tempFilename + "." + format

	if file, err = os.OpenFile(getAbsCacheDirname()+tempFilename, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	switch format {
	case "jpeg":
		if err = jpeg.Encode(file, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			log.Fatal(err)
		}
		break
	case "jpg":
		if err = jpeg.Encode(file, img, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			log.Fatal(err)
		}
		break
	case "png":
		if err = png.Encode(file, img); err != nil {
			log.Fatal(err)
		}
		break
	default:
		log.Fatalf("unknown format %s\n", format)
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
		tempFilename = "/random_image"
		absPath      string
		err          error
		file         *os.File
		randomImage  []byte
		img          image.Image
		format       string
	)

	if randomImage, err = imageGetter.GetRandomImage(); err != nil {
		log.Fatal(err)
	}

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(string(randomImage)))

	if img, format, err = image.Decode(reader); err != nil {
		log.Fatal(err)
	}

	if file, err = os.OpenFile(getAbsCacheDirname()+tempFilename+"."+format, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if err = png.Encode(file, img); err != nil {
		log.Fatal(err)
	}

	if absPath, err = filepath.Abs(getAbsCacheDirname() + tempFilename + "." + format); err != nil {
		log.Fatal(err)
	}

	if err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", absPath).Run(); err != nil {
		log.Fatal(err)
	}
}
