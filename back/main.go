package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/api"
	"github.com/goncharovnikita/wallpaperize/back/crawler"
	"github.com/goncharovnikita/wallpaperize/back/server"
	"github.com/goncharovnikita/wallpaperize/back/utils"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Spec app specification
type Spec struct {
	Port                 int     `required:"false" default:"8080" envconfig:"PORT"`
	BuildsPath           string  `required:"false" default:"uploads" envconfig:"BUILDS_PATH"`
	RandomImagesPath     string  `required:"false" default:"random_images" envconfig:"RANDOM_IMAGES_PATH"`
	MaxRandomDiskUsageGB float64 `required:"false" default:"1" envconfig:"MAX_RANDOM_DISK_USAGE_GB"`
}

func main() {
	godotenv.Load()
	spec := Spec{}
	envconfig.MustProcess("", &spec)

	os.Mkdir(spec.BuildsPath, 0777)
	os.Mkdir(spec.RandomImagesPath, 0777)

	maxRandomUsageInt := int64(0)
	if spec.MaxRandomDiskUsageGB > 5 {
		log.Fatal("we cannot afford that much disk space")
	}
	maxRandomUsageInt = utils.GetBytesFromGigabytes(spec.MaxRandomDiskUsageGB)

	rndCrawler := crawler.NewRandomCrawler(spec.RandomImagesPath, maxRandomUsageInt, api.UnsplashAPI{})
	go rndCrawler.Crawl()

	minCrawer := crawler.NewMinCrawler([]string{spec.RandomImagesPath})
	go minCrawer.Crawl()

	s := server.NewServer(spec.BuildsPath, spec.RandomImagesPath)
	go s.Serve()
	println("app listen on :: ", spec.Port)

	http.ListenAndServe(fmt.Sprintf(":%d", spec.Port), nil)
}
