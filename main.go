package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	unsplashRandomFlag = flag.Bool("u", false, "set random picture from unsplash as wallpaper")
	picsumRandomFlag   = flag.Bool("p", false, "set random picture from lorem picsum as wallpaper")
	daemonizeFlag      = flag.Bool("d", false, "run as daemon")
	cacheDirname       = "/.wallpaperize_cache"
	absCacheDirname    string
	absRandomDirname   string
	ch                 cache
	conf               config
	cl                 cleaner
)

func init() {
	absCacheDirname = getAbsCacheDirname()
	absRandomDirname = absCacheDirname + "/random"
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	createCacheFolder()
	conf.setup()
	ch.cache()
	flag.Parse()
}

func main() {
	switch true {
	case !isGNOMECompatible():
		fmt.Printf("cannot perform wallpaperize - such system is not compatible yet\n")
		break
	case *unsplashRandomFlag != false:
		setRandomPhoto()
		break
	// case *picsumRandomFlag != false:
	// 	var picsumAPI PicsumAPI
	// 	setRandomPhoto(picsumAPI)
	// 	break
	case *daemonizeFlag != false:
		runAsDaemon()
		break
	default:
		setPhotoOfTheDay()
		break
	}
	cl.cleanRandomImages()
}
