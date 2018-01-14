package main

import (
	"log"
	"os"
	"os/user"
	"strings"
	"time"
)

var Desktop = os.Getenv("XDG_CURRENT_DESKTOP")

func createCacheFolder() {
	var (
		err error
		usr *user.User
	)
	if usr, err = user.Current(); err != nil {
		log.Fatal(err)
	}

	if err = os.Mkdir(usr.HomeDir+cacheDirname, 0777); err != nil {
		var ok bool
		if _, ok = err.(*os.PathError); !ok {
			log.Fatal(err)
		}
	}

	if err = os.Mkdir(usr.HomeDir+cacheDirname+"/random", 0777); err != nil {
		var ok bool
		if _, ok = err.(*os.PathError); !ok {
			log.Fatal(err)
		}
	}
}

func getAbsCacheDirname() string {
	var (
		err error
		usr *user.User
	)
	if usr, err = user.Current(); err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir + cacheDirname
}

func isGNOMECompatible() bool {
	return strings.Contains(Desktop, "GNOME") || Desktop == "Unity" || Desktop == "Pantheon"
}

func compareDates(v1 *time.Time, v2 *time.Time) bool {
	return v1.Year() == v2.Year() && v1.Month() == v2.Month() && v1.Day() == v2.Day()
}
