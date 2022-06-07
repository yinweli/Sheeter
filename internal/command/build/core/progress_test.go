package core

import (
	"io/ioutil"
	"testing"

	"github.com/schollz/progressbar/v3"
	"github.com/stretchr/testify/assert"
)

func TestProgress(t *testing.T) {
	progress := mockProgress(1000)
	percent := 0.0

	for i := 0; i <= 1000; i++ {
		progress.Add(1)
		assert.LessOrEqual(t, percent, progress.bar.State().CurrentPercent)
		percent = progress.bar.State().CurrentPercent
	} // for

	assert.True(t, progress.bar.IsFinished())

	progress = mockProgress(1000)
	progress.Finish()
	assert.True(t, progress.bar.IsFinished())
}

func TestNewProgress(t *testing.T) {
	progress := NewProgress(0, "", ioutil.Discard)
	assert.NotNil(t, progress)
}

func mockProgress(max int) *Progress {
	return &Progress{
		bar: progressbar.NewOptions(
			max,
			progressbar.OptionSetWriter(ioutil.Discard),
		),
	}
}
