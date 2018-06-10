package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/goncharovnikita/wallpaperize/app/api"
)

func (a application) Info(format string) {
	switch format {
	case "", "simple":
		a.simpleInfo()
		break
	case "json":
		a.jsonInfo()
		break
	default:
		println("Unsupported format type", format)
	}
}

func (a application) jsonInfo() {
	data := api.AppInfo{
		AppVersion: appVersion,
		Arch:       runtime.GOARCH,
		OS:         runtime.GOOS,
		Build:      appBuild,
	}

	stringed, err := json.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(stringed))
}

func (a application) simpleInfo() {
	fmt.Printf(`
Application version %s

**
%s
**
%s
**

`, appVersion,
		a.info("daily photos storage", a.cache.getDailyPath()),
		a.info("random photos storage", a.cache.getRandomPath()),
	)
}

func (a application) info(prefix string, path string) string {
	inf := a.getDirInfo(path)
	if len(inf) < 1 {
		return fmt.Sprintf("** %s has no images **", prefix)
	}

	sumSize := int64(0)
	for _, v := range inf {
		sumSize += v.Size()
	}

	return fmt.Sprintf(
		"** %s has %d images, summary size is %s **",
		prefix, len(inf), getSizeAsString(sumSize))
}

func (a application) getDirInfo(path string) []os.FileInfo {
	result := make([]os.FileInfo, 0)

	filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".jpg") {
			return nil
		}
		result = append(result, info)
		return nil
	})

	return result
}
