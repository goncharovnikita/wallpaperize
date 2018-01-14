package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	fileperm = os.FileMode(0666)
)

type configRandomPhoto struct {
	Name         string `json:"name"`
	RelativePath string `json:"relative_path"`
}

func (c configRandomPhoto) getPathWithName() string {
	return c.RelativePath + c.Name
}

type configDailyImage struct {
	Type      string `json:"type"`
	Date      string `json:"date"`
	Extension string `json:"extension"`
}

func (c configDailyImage) getName() string {
	return c.Type + "_" + c.Date + "." + c.Extension
}

type configStructure struct {
	DailyImages  map[string]configDailyImage
	RandomPhotos []configRandomPhoto
}

type config struct {
	name      string
	absName   string
	exists    bool
	structure configStructure
}

func (c config) setup() {
	fmt.Println("Init config...")
	conf = config{name: "config.json", absName: absCacheDirname + "/config.json"}
	var err error
	if conf.exists, err = conf.checkExists(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config file exists: %t\n", conf.exists)

	if !conf.exists {
		if err = conf.create(); err != nil {
			log.Fatal(err)
		}
	}
}

// create config file if not exists
func (c config) create() (err error) {
	fmt.Println("creating config file")
	_, err = os.OpenFile(c.absName, os.O_CREATE, fileperm)
	return
}

// get config entity
func (c config) read() (result []byte, err error) {
	fmt.Println("reading config file")
	var (
		file *os.File
	)

	if file, err = os.OpenFile(c.absName, os.O_RDONLY, fileperm); err != nil {
		return
	}

	defer file.Close()

	if result, err = ioutil.ReadAll(file); err != nil {
		return
	}

	return
}

// parse config
// get count of cached images, and ordered
// array of all cached images
func (c config) parseConfig() (result configStructure, err error) {
	fmt.Println("parsing config file")
	var entity []byte
	if entity, err = c.read(); err != nil {
		return
	}
	err = json.Unmarshal(entity, result)
	return
}

// switch to next position and remove previous
func (c config) switchNext() (next string, err error) {
	fmt.Println("switch next config file")
	var (
		struc configStructure
		file  *os.File
		data  []byte
	)
	if struc, err = c.parseConfig(); err != nil {
		return
	}

	if len(struc.RandomPhotos) < 2 {
		return "", nil
	}

	next = struc.RandomPhotos[1].getPathWithName()

	struc.RandomPhotos = struc.RandomPhotos[1:]

	if data, err = json.Marshal(struc); err != nil {
		return
	}

	if file, err = os.OpenFile(c.absName, os.O_RDWR|os.O_TRUNC, fileperm); err != nil {
		return
	}

	defer file.Close()

	if _, err = file.Write(data); err != nil {
		return
	}

	return
}

// addNewToConfig add new name to config
func (c config) addRandomPhoto(name string) (err error) {
	fmt.Printf("adding %s to config file\n", name)
	var (
		file  *os.File
		struc configStructure
		data  []byte
	)

	if struc, err = c.parseConfig(); err != nil {
		log.Print(err)
		return
	}

	struc.RandomPhotos = append(struc.RandomPhotos, configRandomPhoto{Name: name, RelativePath: "/random/"})

	if data, err = json.Marshal(struc); err != nil {
		log.Print(err)
		return
	}

	if file, err = os.OpenFile(c.absName, os.O_RDWR, fileperm); err != nil {
		return
	}

	defer file.Close()

	if _, err = file.Write(data); err != nil {
		return
	}

	return
}

// add daily photo
func (c config) addDailyPhoto(p configDailyImage) (err error) {
	var (
		struc configStructure
		data  []byte
		file  *os.File
	)

	if struc, err = c.parseConfig(); err != nil {
		log.Print(err)
		return
	}

	struc.DailyImages[p.Type] = p

	if data, err = json.Marshal(struc); err != nil {
		log.Print(err)
		return
	}

	if file, err = os.OpenFile(c.absName, os.O_RDWR, fileperm); err != nil {
		return
	}

	defer file.Close()

	if _, err = file.Write(data); err != nil {
		return
	}

	return
}

// checking for config file exists
func (c config) checkExists() (result bool, err error) {
	fmt.Println("checking if config file exists")
	if _, err = os.OpenFile(c.absName, os.O_RDONLY, fileperm); err != nil {
		if _, ok := err.(*os.PathError); ok {
			return false, nil
		}
		return
	}
	result = true

	return
}
