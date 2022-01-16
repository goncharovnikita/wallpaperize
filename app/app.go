package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/kardianos/osext"
	"github.com/reujab/wallpaper"
)

type imageGetter interface {
	GetImage() (string, error)
}

type application struct {
	cache       *cacher
	rec         *recoverer
	dailyGetter imageGetter
	rndGetter   imageGetter
	logger      *log.Logger
}

func newApplication(
	cache *cacher,
	rec *recoverer,
	dailyGetter imageGetter,
	rndGetter imageGetter,
	logger *log.Logger,
) *application {
	return &application{
		cache:       cache,
		rec:         rec,
		dailyGetter: dailyGetter,
		rndGetter:   rndGetter,
		logger:      logger,
	}
}

func (a application) Daily() error {
	name, err := a.dailyGetter.GetImage()
	if err != nil {
		return err
	}

	err = wallpaper.SetFromFile(name)
	if err != nil {
		return err
	}

	return nil
}

func (a application) Info(format string) error {
	switch format {
	case "", "simple":
		return a.simpleInfo()
	case "json":
		return a.jsonInfo()
	default:
		return fmt.Errorf("unsupported error format: %s", format)
	}
}

type appInfo struct {
	AppVersion   string   `json:"app_version"`
	Arch         string   `json:"arch"`
	OS           string   `json:"os"`
	Build        string   `json:"build"`
	DailyImages  []string `json:"daily_images"`
	RandomImages []string `json:"random_images"`
}

func (a application) jsonInfo() error {
	randomDir := a.cache.getRandomPath()
	dailyDir := a.cache.getDailyPath()
	randomInfo := a.getDirInfo(randomDir)
	dailyInfo := a.getDirInfo(dailyDir)
	data := appInfo{
		AppVersion:   appVersion,
		Arch:         runtime.GOARCH,
		OS:           runtime.GOOS,
		Build:        appBuild,
		RandomImages: getFileNames(randomDir, randomInfo),
		DailyImages:  getFileNames(dailyDir, dailyInfo),
	}

	stringed, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(stringed))

	return nil
}

func (a application) simpleInfo() error {
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

	return nil
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

func (a application) GetSelected() error {
	selected, err := wallpaper.Get()
	if err != nil {
		return err
	}

	a.logger.Println(selected)

	return nil
}

func (a application) Place() error {
	filename, err := osext.Executable()
	if err != nil {
		return err
	}

	fmt.Println(filename)

	return nil
}

func (a application) Random(loadOnly bool) error {
	fname, err := a.rndGetter.GetImage()
	if err != nil {
		return err
	}

	if !loadOnly {
		return wallpaper.SetFromFile(fname)
	}

	return nil
}

func (a application) Restore() error {
	if a.rec.failed {
		return fmt.Errorf("could not recover original image")
	}

	return wallpaper.SetFromFile(a.rec.getRecoverFilepath())
}

func (a application) Set(path string) error {
	path = strings.Replace(path, "file://", "", 1)
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		a.setFromRemote(path)

		return nil
	}

	err := wallpaper.SetFromFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (a application) setFromRemote(path string) error {
	parts := strings.Split(path, "/")
	last := parts[len(parts)-1]
	name := a.cache.getRandomPath() + "/" + last

	_, err := os.Stat(name)
	if err != nil {
		resp, err := http.Get(path)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			return err
		}

		defer file.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			return err
		}
	}

	time.Sleep(100 * time.Millisecond)

	err = wallpaper.SetFromFile(name)
	if err != nil {
		return nil
	}

	return nil
}
