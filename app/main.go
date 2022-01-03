package main

import (
	"log"

	"github.com/goncharovnikita/wallpaperize/app/api"
	"github.com/goncharovnikita/wallpaperize/app/cmd"
	"github.com/goncharovnikita/wallpaperize/app/daily"
	"github.com/goncharovnikita/wallpaperize/app/random"
)

var (
	appVersion string
	appBuild   string
)

func main() {
	logger := log.Default()
	logger.SetFlags(log.LstdFlags | log.Lshortfile)

	master := getWallmaster()
	cache := newCacher()
	rec, err := newRecoverer(master, cache.getRecoverPath())
	if err != nil {
		log.Println(err)

		return
	}

	dailyGetter := daily.NewDailyGetter(api.BingAPI{}, cache.getDailyPath())
	rndGetter := random.NewRandomImageGetter(api.UnsplashAPI{}, cache.getRandomPath())

	app := newApplication(
		cache,
		rec,
		master,
		dailyGetter,
		rndGetter,
		logger,
	)

	cmd.Execute(app)
}
