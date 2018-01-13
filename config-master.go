package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	conf     config
	fileperm = os.FileMode(0666)
)

type config struct {
	name    string
	absName string
	exists  bool
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Init config...")
	absCacheDirname = getAbsCacheDirname()
	conf = config{name: "config", absName: absCacheDirname + "/config"}
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
	_, err = os.OpenFile(c.absName, os.O_CREATE, fileperm)
	return
}

// get config entity
func (c config) read() (result string, err error) {
	var (
		file   *os.File
		entity []byte
	)

	if file, err = os.OpenFile(c.absName, os.O_RDONLY, fileperm); err != nil {
		return
	}

	defer file.Close()

	if entity, err = ioutil.ReadAll(file); err != nil {
		return
	}

	return string(entity), nil
}

// parse config
// get count of cached images, and ordered
// array of all cached images
func (c config) parseConfig() (names []string, err error) {
	var entity string
	if entity, err = c.read(); err != nil {
		return
	}
	names = strings.Split(entity, "\n")
	return
}

// switch to next position and remove previous
func (c config) switchNext() (next string, err error) {
	var (
		names []string
		file  *os.File
	)
	if names, err = c.parseConfig(); err != nil {
		return
	}

	if len(names) < 2 {
		return "", nil
	}

	next = names[1]

	if file, err = os.OpenFile(c.absName, os.O_RDWR, fileperm); err != nil {
		return
	}

	defer file.Close()

	if _, err = file.Write([]byte(strings.Join(names[1:], "\n"))); err != nil {
		return
	}

	return
}

// addNewToConfig add new name to config
func (c config) add(name string) (err error) {
	var (
		file *os.File
	)

	name = name + "\n"

	if file, err = os.OpenFile(c.absName, os.O_APPEND|os.O_RDWR, fileperm); err != nil {
		return
	}

	defer file.Close()

	if _, err = file.Write([]byte(name)); err != nil {
		return
	}

	return
}

// checking for config file exists
func (c config) checkExists() (result bool, err error) {

	if _, err = os.OpenFile(c.absName, os.O_RDONLY, fileperm); err != nil {
		if _, ok := err.(*os.PathError); ok {
			return false, nil
		}
		return
	}
	result = true

	return
}
