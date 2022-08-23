package iox

import (
	"os"
)

// Check wether a path exists or not.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
