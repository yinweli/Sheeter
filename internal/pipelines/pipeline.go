package pipelines

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Executor 執行函式類型
type Executor = func(material any) error

// Execute 執行管線
func Execute(name string, material []any, executor []Executor) []error {
	errs := []error{}

	if len(material) == 0 || len(executor) == 0 {
		return errs
	} // if

	count := len(material) * len(executor)
	signaler := utils.NewWaitGroup(count)
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(signaler))
	progressbar := progress.AddBar(
		int64(count),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name(name),
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

				signaler.Done()
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
