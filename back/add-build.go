package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/goncharovnikita/wallpaperize/app/cerrors"

	wallpaperize "github.com/goncharovnikita/wallpaperize/app/api"
)

func addBuild(path string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		err := handleNewBuild(r.Body, path)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte("error occured"))
			return
		}
		rw.WriteHeader(200)
	}
}

func handleNewBuild(rdr io.Reader, path string) error {
	tmpName := os.TempDir() + "/" + randStringBytes(10)
	tmpFile, err := os.OpenFile(tmpName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return err
	}

	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, rdr)
	if err != nil {
		log.Println(err)
		return err
	}

	info := bytes.NewBuffer([]byte{})

	cmd := exec.Command(tmpName, "info", "-o", "json")
	cmd.Stdout = info
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		return err
	}

	var serializedInfo wallpaperize.AppInfo

	err = json.Unmarshal(info.Bytes(), &serializedInfo)
	if err != nil {
		log.Println(err)
		return err
	}

	fname := path + "/" + strings.Join(
		[]string{
			serializedInfo.OS,
			serializedInfo.Arch,
			serializedInfo.AppVersion,
		}, "-",
	)

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
