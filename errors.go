package main

import (
	"fmt"
)

type statNoSuchFileErr struct {
	message string
}

func newStatNoSuchFileErr(msg string) statNoSuchFileErr {
	return statNoSuchFileErr{
		message: msg,
	}
}

func (s statNoSuchFileErr) Error() string {
	return fmt.Sprintf("stat %s: no such file or directory", s.message)
}
