package util

import (
	"fmt"
	"io"
	"time"

	"Sheeter/internal"

	"github.com/schollz/progressbar/v3"
)

// NewProgressBar 建立進度條物件
func NewProgressBar(max int64, desc string, writer io.Writer) *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		max,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(writer),
		progressbar.OptionSetWidth(40),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n%s complete!\n", desc)
		}),
		progressbar.OptionSpinnerType(33),
	)
	_ = bar.RenderBlank()

	return bar
}

// CalcProgress 計算進度值
func CalcProgress(count int, task int) int {
	return int(float32(count*task) * internal.ProgressFactor)
}
