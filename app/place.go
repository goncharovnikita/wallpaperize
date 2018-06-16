package main

import (
	"fmt"

	"github.com/kardianos/osext"
)

// Place implementatipn
func (a application) Place() {
	filename, _ := osext.Executable()
	fmt.Println(filename)
}
