package main

import (
	"log"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/api"
	"github.com/goncharovnikita/wallpaperize/app/cmd"
	"github.com/goncharovnikita/wallpaperize/app/daily"
	"github.com/goncharovnikita/wallpaperize/app/random"
	"github.com/goncharovnikita/wallpaperize/app/store"
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
	tokenStore := store.NewFile(cache.getConfigPath())
	clientID := os.Getenv("WALLPAPERIZE_UNSPLASH_APP_ID")
	clientSecret := os.Getenv("WALLPAPERIZE_UNSPLASH_APP_SECRET")
	unsplashAuthorizer := api.NewUnsplashAuthorizer(
		clientID,
		clientSecret,
		tokenStore,
	)
	unsplashApi := api.NewUnsplashAPI(unsplashAuthorizer)
	rndGetter := random.NewRandomImageGetter(unsplashApi, cache.getRandomPath())

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
