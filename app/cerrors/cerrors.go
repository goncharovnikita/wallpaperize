package cerrors

import (
	"fmt"
)

type StatNoSuchFileErr struct {
	message string
}

func NewStatNoSuchFileErr(msg string) StatNoSuchFileErr {
	return StatNoSuchFileErr{
		message: msg,
	}
}

func (s StatNoSuchFileErr) Error() string {
	return fmt.Sprintf("stat %s: no such file or directory", s.message)
}
