package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgress(t *testing.T) {
	progress := NewProgress(1000, "", os.Stdout)
	percent := 0.0

	for i := 0; i <= 1000; i++ {
		progress.Add(1)
		assert.LessOrEqual(t, percent, progress.bar.State().CurrentPercent)
		percent = progress.bar.State().CurrentPercent
	} // for

	assert.True(t, progress.bar.IsFinished())

	progress = NewProgress(1000, "", ioutil.Discard)
	progress.Finish()
	assert.True(t, progress.bar.IsFinished())
}

func TestNewProgress(t *testing.T) {
	progress := NewProgress(0, "", ioutil.Discard)
	assert.NotNil(t, progress)
}
