package iox

import (
	"github.com/yookoala/realpath"
)

func Realpath(path string) string {
	rPath, err := realpath.Realpath(path)
	if err != nil {
		return path
	} else {
		return rPath
	}
}
