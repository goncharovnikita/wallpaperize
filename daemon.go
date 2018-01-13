package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// runAsDaemon start infinite loop
// with changing wallpaper every n minutes
func runAsDaemon() {
	fmt.Printf("Starting as daemon...\n")
	signals := make(chan os.Signal, 1)

	signal.Notify(signals)

	go func() {
		s := <-signals
		fmt.Printf("\nGet %s signal!\nStopping...\n", s)
		daemonCleanup()
		os.Exit(1)
	}()

	for {
		seconds := rand.Int31n(15)
		fmt.Printf("Next change will be in a %ds\n", seconds)
		time.Sleep(time.Second * time.Duration(seconds))
		setRandomPhoto()
	}
}

// daemonCleanup performs cleaning operations
func daemonCleanup() {

}
