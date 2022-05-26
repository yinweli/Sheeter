package util

import (
	"fmt"
	"io"
	"time"

	"github.com/schollz/progressbar/v3"
)

// NewProgressBar 建立進度條物件
func NewProgressBar(max int, desc string, writer io.Writer) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(
		max,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(writer),
		progressbar.OptionSetWidth(40),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\r")
		}),
		progressbar.OptionSpinnerType(14),
	)
	_ = bar.RenderBlank()

	return bar
}

// CalcProgress 計算進度值
func CalcProgress(cellSize int, taskCount int) int {
	return cellSize * taskCount
}
