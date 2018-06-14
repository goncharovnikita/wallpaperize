package main

import "log"

// Random implementation
func (a application) Random(loadOnly bool) {
	fname, err := a.rndGetter.GetImage()
	if err != nil {
		log.Fatal(err)
	}

	if !loadOnly {
		a.master.SetFromFile(fname)
	}
}
