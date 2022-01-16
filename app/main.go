package main

import (
	"log"

	"github.com/goncharovnikita/wallpaperize/app/api"
	"github.com/goncharovnikita/wallpaperize/app/cmd"
	"github.com/goncharovnikita/wallpaperize/app/daily"
	"github.com/goncharovnikita/wallpaperize/app/random"
	"github.com/goncharovnikita/wallpaperize/back/client"
)

var (
	appVersion string
	appBuild   string
)

func main() {
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)

	cache := newCacher()
	rec, err := newRecoverer(cache.getRecoverPath())
	if err != nil {
		log.Println(err)

		return
	}

	dailyGetter := daily.NewDailyGetter(api.BingAPI{}, cache.getDailyPath())
	imagesApi := client.NewHTTP("http://goncharovnikita.com/wallpaperize/api")
	rndGetter := random.NewRandomImageGetter(imagesApi, cache.getRandomPath())

	app := newApplication(
		cache,
		rec,
		dailyGetter,
		rndGetter,
		logger,
	)

	cmd.Execute(app)
}
