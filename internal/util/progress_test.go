package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgressBar(t *testing.T) {
	progress := NewProgressBar("desc", os.Stdout)

	assert.NotNil(t, progress)
	assert.Nil(t, progress.Finish())
}

func TestCalcProgress(t *testing.T) {
	assert.Equal(t, int(float32(5*7)*progressFactor), CalcProgress(5, 7))
}
