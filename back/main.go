package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/goncharovnikita/wallpaperize/app/api"
	"github.com/goncharovnikita/wallpaperize/back/crawler"
)

const (
	VERSION_HEADER = "BUILD_VERSION"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	path := os.Getenv("BUILDS_PATH")
	if path == "" {
		path = "uploads"
		os.Mkdir(path, 0777)
	}

	randomImagesPath := os.Getenv("RANDOM_IMAGES_PATH")
	if randomImagesPath == "" {
		randomImagesPath = "random_images"
		os.Mkdir(randomImagesPath, 0777)
	}

	maxRandomUsageInt := int64(0)
	maxRandomUsage := os.Getenv("MAX_RANDOM_DISK_USAGE")
	if maxRandomUsage == "" {
		maxRandomUsageInt = getBytesFromGigabytes(1)
	} else {
		gbInt, err := strconv.Atoi(maxRandomUsage)
		if err != nil {
			log.Fatal(err)
		}
		if gbInt > 5 {
			log.Println("we cannot afford that")
		}
		maxRandomUsageInt = getBytesFromGigabytes(gbInt)
	}

	rndCrawler := crawler.NewRandomCrawler(randomImagesPath, maxRandomUsageInt, api.UnsplashAPI{})
	rndCrawler.Crawl()

	serve(path)
	println("app listen on :: ", port)

	http.ListenAndServe(":"+port, nil)
}
