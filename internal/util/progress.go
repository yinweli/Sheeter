package util

import (
	"fmt"
	"io"

	"github.com/schollz/progressbar/v3"
)

// NewProgress 建立進度條
func NewProgress(max int64, desc string, writer io.Writer) *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		max,
		progressbar.OptionSetDescription(fmt.Sprintf("%-20s", desc)),
		progressbar.OptionSetWriter(writer),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		// progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n%s complete!\n", desc)
		}),
		progressbar.OptionSpinnerType(14),
	)
	_ = bar.RenderBlank()

	return bar
}
