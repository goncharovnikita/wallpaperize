package main

import (
	"log"
	"os"
	"path/filepath"
)

type cacher struct {
	dir string
}

func newCacher() *cacher {
	result := &cacher{}
	result.setCacheDir()
	result.initCacheDir()
	return result
}

func (c cacher) saveDaily(data []byte) string {
	fname := c.dir + "/daily/daily.jpg"
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	return fname
}

func (c cacher) getRecoverPath() string {
	return c.dir + "/preserved"
}

func (c cacher) getDailyPath() string {
	return c.dir + "/daily"
}

func (c cacher) getRandomPath() string {
	return c.dir + "/random"
}

func (c cacher) getConfigPath() string {
	return c.dir + "/config/config.txt"
}

func (c *cacher) setCacheDir() {
	result, err := filepath.Abs(os.Getenv("HOME") + "/.wallpaperize_cache")
	if err != nil {
		log.Fatal(err)
	}

	c.dir = result
}

func (c *cacher) initCacheDir() {
	ensureDir(c.dir)
	ensureDir(c.dir + "/daily")
	ensureDir(c.dir + "/random")
	ensureDir(c.dir + "/preserved")
	ensureDir(c.dir + "/config")
}
