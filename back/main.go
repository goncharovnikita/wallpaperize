package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	Debug                bool    `required:"false" default:"true"`
}

func main() {
	godotenv.Load()
	spec := Spec{}
	envconfig.MustProcess("", &spec)

	os.Mkdir(spec.BuildsPath, 0777)
	os.Mkdir(spec.RandomImagesPath, 0777)

	s := server.NewServer(spec.BuildsPath, spec.RandomImagesPath, spec.Debug)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		srv := http.Server{
			Addr:    fmt.Sprintf(":%d", spec.Port),
			Handler: s,
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server listen on :%d\n", spec.Port)

	<-done
	log.Print("Server Stopped")
}

func getMaxDiskUsage(spec Spec) int64 {
	maxRandomUsageInt := int64(0)
	if spec.MaxRandomDiskUsageGB > 5 {
		log.Fatal("we cannot afford that much disk space")
	}
	maxRandomUsageInt = utils.GetBytesFromGigabytes(spec.MaxRandomDiskUsageGB)

	return maxRandomUsageInt
}
