package iox

import "os"

// Check wether a directory exists or not.
func DirExists(f string) bool {
	fileInfo, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir()
}
