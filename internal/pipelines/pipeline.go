package pipelines

import (
	"fmt"
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
)

// Executor 執行函式類型
type Executor = func(material any) error

// Execute 執行管線; 說是管線機制, 但其實更像帶有進度條顯示的重複執行功能;
// 執行時會用executor列表中的執行函式對每個material執行一次;
// 不管有無錯誤, 都會把所有該執行的都跑完, 並回傳執行的錯誤列表
func Execute(name string, material []any, executor []Executor) []error {
	errs := []error{}

	if len(material) == 0 || len(executor) == 0 {
		return errs
	} // if

	count := len(material) * len(executor)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(count)
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(waitGroup))
	progressbar := progress.AddBar(
		int64(count),
		mpb.AppendDecorators(
			decor.Name(fmt.Sprintf("%-10s ", name)),
			decor.CountersNoUnit("(%6d/%6d) ", decor.WCSyncWidth),
			decor.NewPercentage("%d "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)
	errors := make(chan error, count)

	for _, itor := range material {
		meta := itor // 多執行緒需要使用中間變數

		go func() {
			for _, itor := range executor {
				if err := itor(meta); err != nil {
					errors <- fmt.Errorf("%w", err)
				} // if

				waitGroup.Done()
				progressbar.Increment()
			} // for
		}()
	} // for

	progress.Wait()
	close(errors) // 先關閉錯誤通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return errs
}
