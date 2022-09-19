package builds

import (
	"sync"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
)

// Encoding 產生編碼資料
func Encoding(runtime *Runtime) (errs []error) {
	const tasks = 1 // 要執行的工作數量

	count := len(runtime.Sector)
	errors := make(chan error) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := sync.WaitGroup{}
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(&signaler))
	progressbar := progress.AddBar(
		int64(count*tasks),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("encoding "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	signaler.Add(count)

	for _, itor := range runtime.Sector {
		runtimeSector := itor // 多執行緒需要使用中間變數

		go func() {
			defer signaler.Done()
			defer runtimeSector.Close()

			if err := encodingJson(runtimeSector); err != nil {
				errors <- err
				return
			} // if

			progressbar.Increment()
		}()
	} // for

	progress.Wait()
	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return errs
}
