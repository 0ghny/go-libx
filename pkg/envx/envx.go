package envx

import (
	"fmt"
	"os"
	"strings"
)

func IsDefined(name string) bool {
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, name) {
			return true
		}
	}
	return false
}

func Find(name string) string {
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, name) {
			return v
		}
	}
	return ""
}

func Print() {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

func GetOrDefault(name string, fallback string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return fallback
}
