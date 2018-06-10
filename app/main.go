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

type application struct {
	cache       *cacher
	rec         *recoverer
	master      Wallmaster
	dailyGetter *daily.Daily
	rndGetter   *random.Random
}

func newApplication() *application {
	master := getWallmaster()
	cache := newCacher()
	rec := newRecoverer(master, cache.getRecoverPath())
	dailyGetter := daily.NewDailyGetter(api.BingAPI{}, cache.getDailyPath())
	rndGetter := random.NewRandomImageGetter(api.UnsplashAPI{}, cache.getRandomPath())
	return &application{
		cache:       cache,
		rec:         rec,
		master:      master,
		dailyGetter: dailyGetter,
		rndGetter:   rndGetter,
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	app := newApplication()
	cmd.Execute(app)
}
