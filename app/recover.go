package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/api"
)

type recoverer struct {
	filepath string
	failed   bool
}

func (r *recoverer) getRecoverFilepath() string {
	return r.filepath
}

func newRecoverer(master api.Wallmaster, path string) (*recoverer, error) {
	if err := ensureFile(path + "/config.txt"); err != nil {
		return nil, err
	}

	rec := &recoverer{}

	if err := rec.initRecoverImage(master, path); err != nil {
		return nil, err
	}

	return rec, nil
}

func (r *recoverer) initRecoverImage(master api.Wallmaster, path string) error {
	file, err := os.OpenFile(path+"/config.txt", os.O_RDWR, 0777)
	if err != nil {
		return fmt.Errorf("error opening config file %s: %w", path+"/config.txt", err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if len(data) < 1 {
		fname, err := master.Get()
		if err != nil {
			r.failed = true

			return nil
		}

		_, err = file.WriteString(fname)
		if err != nil {
			return err
		}

		r.filepath = fname

		return nil
	}

	r.filepath = string(data)

	return nil
}
