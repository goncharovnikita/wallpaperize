package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func (a application) Set(path string) {
	path = strings.Replace(path, "file://", "", 1)
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		a.setFromRemote(path)
		return
	}
	err := a.master.SetFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	println("** success **")
}

func (a application) setFromRemote(path string) {
	parts := strings.Split(path, "/")
	last := parts[len(parts)-1]
	name := a.cache.getRandomPath() + "/" + last
	_, err := os.Stat(name)
	if err != nil {
		resp, err := http.Get(path)
		if err != nil {
			println("ERROR")
			log.Fatal(err)
		}

		defer resp.Body.Close()

		file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			println("ERROR")
			log.Fatal(err)
		}

		io.Copy(file, resp.Body)
		file.Close()
	}

	time.Sleep(100 * time.Millisecond)

	println(name)

	err = a.master.SetFromFile(name)
	if err != nil {
		log.Fatal(err)
	}
	println("** success **")
}
