package main

import (
	"log"
	"os"
	"os/user"
	"strings"
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
