package main

import (
	"log"

	"github.com/goncharovnikita/wallpaperize/api"
	"github.com/goncharovnikita/wallpaperize/daily"
)

var app *application

type application struct {
	cache       *cacher
	rec         *recoverer
	master      Wallmaster
	dailyGetter ImageGetter
}

func newApplication() *application {
	master := getWallmaster()
	cache := newCacher()
	rec := newRecoverer(master, cache.getRecoverPath())
	dailyGetter := daily.NewDailyGetter(api.BingAPI{})
	return &application{
		cache:       cache,
		rec:         rec,
		master:      master,
		dailyGetter: dailyGetter,
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app = newApplication()
}

func main() {
	dail, err := app.dailyGetter.GetImage()
	if err != nil {
		log.Fatal(err)
	}

	name := app.cache.saveDaily(dail)
	app.master.SetFromFile(name)
}
