package main

import "os"

// Informator provide information about images path
type Informator interface {
	GetInfo() []os.FileInfo
}
