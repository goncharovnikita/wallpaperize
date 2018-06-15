package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/api"
)

type recoverer struct {
	filepath string
}

func (r *recoverer) getRecoverFilepath() string {
	return r.filepath
}

func newRecoverer(master api.Wallmaster, path string) *recoverer {
	ensureFile(path + "/config.txt")
	rec := &recoverer{}
	rec.initRecoverImage(master, path)
	return rec
}

func (r *recoverer) initRecoverImage(master api.Wallmaster, path string) {

	file, err := os.OpenFile(path+"/config.txt", os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if len(data) < 1 {
		fname, err := master.Get()
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.WriteString(fname)
		if err != nil {
			log.Fatal(err)
		}

		r.filepath = fname
		return
	}

	r.filepath = string(data)
}
