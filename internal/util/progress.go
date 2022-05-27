package util

import (
	"fmt"
	"io"
	"time"

	"github.com/schollz/progressbar/v3"
)

const progressDefault = 99999999 // 預設進度值, 先給個很大的進度值, 後面如果有變更再通過ProgressBar.ChangeMax去改
const progressFactor = 1.2       // 進度值乘數, 多了0.2是要包括可能會出現的額外工作量

// NewProgressBar 建立進度條物件, 如果不想要顯示進度條(例如單元測試時), 可以用ioutil.Discard填入writer參數中
func NewProgressBar(desc string, writer io.Writer) *progressbar.ProgressBar {
	bar := progressbar.NewOptions64(
		progressDefault,
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
	return int(float32(count*task) * progressFactor)
}
