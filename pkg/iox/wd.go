package iox

import (
	"os"
	"path/filepath"
)

func GetBaseWd() string {
	currDir, err := os.Getwd()
	if err != nil {
		currDir = "./"
	}
	return filepath.Base(currDir)
}

func Getwd() string {
	currDir, err := os.Getwd()
	if err != nil {
		return "./"
	}
	return currDir
}

func GetAbsWd() (string, error) {
	absPath, err := filepath.Abs(Getwd())
	if err != nil {
		return "", err
	}
	return absPath, nil
}
