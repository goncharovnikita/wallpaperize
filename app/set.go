package main

import "log"

func (a application) Set(path string) {
	err := a.master.SetFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	println("** success **")
}
