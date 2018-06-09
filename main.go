package main

import (
	"log"

	"github.com/goncharovnikita/wallpaperize/api"
	"github.com/goncharovnikita/wallpaperize/cmd"
	"github.com/goncharovnikita/wallpaperize/daily"
)

const (
	appVersion = "1.0.0"
)

type application struct {
	cache       *cacher
	rec         *recoverer
	master      Wallmaster
	dailyGetter *daily.Daily
}

func newApplication() *application {
	master := getWallmaster()
	cache := newCacher()
	rec := newRecoverer(master, cache.getRecoverPath())
	dailyGetter := daily.NewDailyGetter(api.BingAPI{}, cache.getDailyPath())
	return &application{
		cache:       cache,
		rec:         rec,
		master:      master,
		dailyGetter: dailyGetter,
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	app := newApplication()
	cmd.Execute(app)
}
