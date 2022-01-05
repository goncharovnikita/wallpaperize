package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/goncharovnikita/wallpaperize/back/repo"
	"github.com/goncharovnikita/wallpaperize/back/server"
	"github.com/goncharovnikita/wallpaperize/back/service"
	"github.com/goncharovnikita/wallpaperize/back/updater"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Spec app specification
type Spec struct {
	Port                 int     `required:"false" default:"8080" envconfig:"PORT"`
	BuildsPath           string  `required:"false" default:"uploads" envconfig:"BUILDS_PATH"`
	RandomImagesPath     string  `required:"false" default:"random_images" envconfig:"RANDOM_IMAGES_PATH"`
	SQLiteStorePath      string  `required:"false" default:"./data/sqlite" envconfig:"SQLITE_STORE_PATH"`
	MaxRandomDiskUsageGB float64 `required:"false" default:"1" envconfig:"MAX_RANDOM_DISK_USAGE_GB"`
	Debug                bool    `required:"false" default:"true"`
	UnsplashAccessToken  string  `required:"true" envconfig:"UNSPLASH_ACCESS_TOKEN"`
	EnableUpdater        bool    `required:"false" envconfig:"ENABLE_UPDATER"`
}

func main() {
	godotenv.Load()
	spec := Spec{}
	envconfig.MustProcess("", &spec)

	infoLogger := log.New(os.Stdout, "INFO ", log.LUTC)
	errLogger := log.New(os.Stderr, "ERR ", log.LUTC)
	updaterLogger := log.New(os.Stdout, "INFO [Updater] ", log.LUTC)
	serverLogger := log.New(os.Stdout, "INFO [Server] ", log.LUTC)

	infoLogger.Println("logger inited")

	os.Mkdir(spec.BuildsPath, 0777)
	os.Mkdir(spec.RandomImagesPath, 0777)
	os.MkdirAll(spec.SQLiteStorePath, 0777)

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/images.db", spec.SQLiteStorePath))
	if err != nil {
		errLogger.Println("failed to create db", err)

		return
	}

	defer db.Close()

	sqliteRepo := repo.NewSQLite(db)

	if err := sqliteRepo.Prepare(); err != nil {
		errLogger.Println("error prepare sqlite repo", err)

		return
	}

	imagesGetter := service.NewImagesGetter(sqliteRepo)

	s := server.NewServer(
		spec.BuildsPath,
		spec.RandomImagesPath,
		imagesGetter,
		serverLogger,
		spec.Debug,
	)

	imagesSetter := service.NewImagesSetter(sqliteRepo)
	updater := updater.NewUnsplash(
		spec.UnsplashAccessToken,
		imagesSetter,
		sqliteRepo,
		updaterLogger,
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	shutdownHTTP := make(chan struct{}, 1)
	shutdownUpdater := make(chan struct{}, 1)

	var shutdownWG sync.WaitGroup

	shutdownWG.Add(2)

	go func(shutdown <-chan struct{}, wg *sync.WaitGroup) {
		defer wg.Done()

		srv := http.Server{
			Addr:    fmt.Sprintf(":%d", spec.Port),
			Handler: s.Listen(),
		}

		go func() {
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				errLogger.Printf("failed to listen: %s\n", err)

				return
			}
		}()

		infoLogger.Printf("http server listening on :%d\n", spec.Port)

		<-shutdown

		infoLogger.Println("shutting down http server...")

		if err := srv.Shutdown(context.Background()); err != nil {
			errLogger.Println("error stopping http server: %w", err)

			return
		}

		infoLogger.Println("http server stopped")
	}(shutdownHTTP, &shutdownWG)

	go func(shutdown <-chan struct{}, wg *sync.WaitGroup) {
		defer wg.Done()

		if !spec.EnableUpdater {
			return
		}

		go func() {
			if err := updater.Run(); err != nil {
				errLogger.Printf("error running updater: %v\n", err)

				return
			}
		}()

		infoLogger.Println("upater is running")

		<-shutdown

		infoLogger.Println("stopping updater...")

		updater.Stop()

		infoLogger.Println("updater is stopped")
	}(shutdownUpdater, &shutdownWG)

	<-done

	infoLogger.Println("stopping application...")

	shutdownHTTP <- struct{}{}
	shutdownUpdater <- struct{}{}

	shutdownWG.Wait()

	infoLogger.Println("appication stopped")
}
