package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/goncharovnikita/wallpaperize/app/api"
	"github.com/goncharovnikita/wallpaperize/back/crawler"
	"github.com/goncharovnikita/wallpaperize/back/server"
	"github.com/goncharovnikita/wallpaperize/back/utils"
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
		maxRandomUsageInt = utils.GetBytesFromGigabytes(1)
	} else {
		gbInt, err := strconv.ParseFloat(maxRandomUsage, 64)
		if err != nil {
			log.Fatal(err)
		}
		if gbInt > 5 {
			log.Fatal("we cannot afford that")
		}
		maxRandomUsageInt = utils.GetBytesFromGigabytes(gbInt)
	}

	rndCrawler := crawler.NewRandomCrawler(randomImagesPath, maxRandomUsageInt, api.UnsplashAPI{})
	go rndCrawler.Crawl()

	minCrawer := crawler.NewMinCrawler([]string{randomImagesPath})
	go minCrawer.Crawl()

	s := server.NewServer(path, randomImagesPath)
	go s.Serve()
	println("app listen on :: ", port)

	http.ListenAndServe(":"+port, nil)
}
