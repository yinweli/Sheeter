package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgressBar(t *testing.T) {
	assert.NotNil(t, NewProgressBar(10, "desc", os.Stdout), "new progressbar failed")
}
