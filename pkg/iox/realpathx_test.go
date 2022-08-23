package iox

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealPathFromCurrentWd(t *testing.T) {
	wd, _ := os.Getwd()
	theRealPath := Realpath(wd)

	assert.NotNil(t, theRealPath)
}
