package main

import (
	"log"
	"runtime"

	"github.com/goncharovnikita/wallpaperize/darwin"
)

func getWallmaster() Wallmaster {
	switch runtime.GOOS {
	case "darwin":
		return darwin.DarwinWallMaster{}
	default:
		log.Fatal("cannot perform wallpaperize - such system is not compatible yet\n")
		return nil
	}
}
