package main

import "log"

// Daily is cmd app daily method implementation
func (a application) Daily() {
	name, err := a.dailyGetter.GetImage()
	if err != nil {
		log.Fatal(err)
	}

	err = a.master.SetFromFile(name)
	if err != nil {
		log.Fatal(err)
	}

	println("** success! **")
}
