package main

import "log"

// GetSelected implementation
func (a application) GetSelected() {
	selected, err := a.master.Get()
	if err != nil {
		println("ERROR")
		log.Println(err)
		return
	}

	println("SUCCESS")
	println(selected)
}
