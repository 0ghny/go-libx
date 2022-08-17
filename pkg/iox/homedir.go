package iox

import (
	"os"
	"path/filepath"
)

// Returns the application home directory inside users home
// or current directory
func GetAppHomedir(appname string) string {
	appHomeDir, err := os.UserHomeDir()
	if err != nil {
		appHomeDir = "./"
	}
	return filepath.Join(appHomeDir, appname)
}
