package main

import (
	"flag"
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
	wallmaster         WallMaster
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
	wallmaster = getOS()
	switch true {
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
