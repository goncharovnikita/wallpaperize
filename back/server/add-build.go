package server

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/cerrors"
	"github.com/goncharovnikita/wallpaperize/back/utils"
)

func addBuild(path string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		name := r.Header.Get(VERSION_HEADER)
		err := handleNewBuild(r.Body, path, name)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("error occured"))
			return
		}
		rw.WriteHeader(200)
	}
}

func handleNewBuild(rdr io.Reader, path string, name string) error {
	tmpName := os.TempDir() + "/" + utils.RandStringBytes(10)
	tmpFile, err := os.OpenFile(tmpName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = io.Copy(tmpFile, rdr)
	if err != nil {
		log.Println(err)
		return err
	}

	tmpFile.Close()

	fname := path + "/" + name

	_, err = os.Stat(fname)
	if err != nil {
		if err.Error() == cerrors.NewStatNoSuchFileErr(fname).Error() {
			file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				log.Println(err)
				return err
			}

			defer file.Close()

			_tmpFile, err := os.Open(tmpName)
			if err != nil {
				log.Println(err)
				return err
			}

			defer _tmpFile.Close()

			_, err = io.Copy(file, _tmpFile)
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			log.Println(err)
			return err
		}
	}

	os.Remove(tmpName)
	return nil
}
