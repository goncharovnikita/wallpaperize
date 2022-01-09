package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"

	"github.com/goncharovnikita/wallpaperize/back/client"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	errLogger := log.New(os.Stderr, "ERR", 0)

	c := client.NewHTTP("http://goncharovnikita.com/wallpaperize/api")

	img, err := c.GetRandomImages(1)
	if err != nil {
		errLogger.Printf("could not get images from api %v", err)

		return
	}

	imgURI, err := storage.ParseURI(img[0].URLs.Small)
	if err != nil {
		errLogger.Printf("error parsing uri: %v", err)

		return
	}

	image := canvas.NewImageFromURI(imgURI)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
}
