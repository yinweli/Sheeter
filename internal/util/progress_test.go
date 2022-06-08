package util

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgress(t *testing.T) {
	progress := NewProgress(0, "", ioutil.Discard)
	assert.NotNil(t, progress)
}
