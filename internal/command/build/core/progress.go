package core

import (
	"fmt"
	"io"
	"time"

	"github.com/schollz/progressbar/v3"
)

// Progress 進度
type Progress struct {
	bar         *progressbar.ProgressBar // 進度條物件
	lastPercent float64                  // 上次進度百分比
}

// Add 增加進度
func (this *Progress) Add(num int) {
	_ = this.bar.Add(num)
	currPercent := this.bar.State().CurrentPercent

	if currPercent-this.lastPercent >= 1.0 {
		this.lastPercent = currPercent
		time.Sleep(10 * time.Millisecond) // 睡眠一下, 讓系統有時間去畫進度條
	} // if
}

// Finish 結束進度
func (this *Progress) Finish() {
	_ = this.bar.Finish()
}

// NewProgress 建立進度物件, 如果不想要顯示進度條(例如單元測試時), 可以用ioutil.Discard填入writer參數中
func NewProgress(max int, desc string, writer io.Writer) *Progress {
	bar := progressbar.NewOptions(
		max,
		progressbar.OptionSetDescription(fmt.Sprintf("%-20s", desc)),
		progressbar.OptionSetWriter(writer),
		progressbar.OptionSetWidth(20),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n%s complete!\n", desc)
		}),
	)
	_ = bar.RenderBlank()

	return &Progress{bar: bar}
}

// NewProgressBytes 建立位元組進度條
func NewProgressBytes(max int64, desc string, writer io.Writer) *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		max,
		progressbar.OptionSetDescription(fmt.Sprintf("%-20s", desc)),
		progressbar.OptionSetWriter(writer),
		progressbar.OptionSetWidth(20),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowBytes(true),
	)
	_ = bar.RenderBlank()

	return bar
}

// NewProgressCount 建立計數進度條
func NewProgressCount(max int64, desc string, writer io.Writer) *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		max,
		progressbar.OptionSetDescription(fmt.Sprintf("%-20s", desc)),
		progressbar.OptionSetWriter(writer),
		progressbar.OptionSetWidth(20),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			fmt.Printf("\n%s complete!\n", desc)
		}),
	)
	_ = bar.RenderBlank()

	return bar
}
