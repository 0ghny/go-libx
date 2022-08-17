package iox

import (
	"os"
	"path/filepath"
)

// Returns current working directory OR ./
func Getwd() string {
	currDir, err := os.Getwd()
	if err != nil {
		currDir = "./"
	}
	return filepath.Base(currDir)
}
