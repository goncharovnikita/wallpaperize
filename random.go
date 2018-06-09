package main

import "log"

// Random implementation
func (a application) Random() {
	fname, err := a.rndGetter.GetImage()
	if err != nil {
		log.Fatal(err)
	}

	a.master.SetFromFile(fname)
}
