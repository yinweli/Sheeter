package builds

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Initialize 初始化
func Initialize(context *Context) (errs []error) {
	for _, itor := range context.Elements {
		context.Sector = append(context.Sector, &ContextSector{
			Element: itor, // 這裡複製會比取用指標好
		})
	} // for

	itemCount := len(context.Sector)
	errors := make(chan error, itemCount) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := utils.NewWaitGroup(itemCount)
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(signaler))
	progressbar := progress.AddBar(
		int64(itemCount),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("initialize "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	for _, itor := range context.Sector {
		ref := itor // 多執行緒需要使用中間變數

		go func() {
			defer signaler.Done()

			if err := initializeSector(context, ref); err != nil {
				errors <- fmt.Errorf("initialize failed: %w", err)
			} // if

			progressbar.Increment()
		}()
	} // for

	progress.Wait()

	if err := initializeStruct(context); err != nil {
		errors <- err
	} // if

	close(errors) // 先關閉結果通訊通道, 免得下面的for變成無限迴圈

	for itor := range errors {
		if itor != nil {
			errs = append(errs, itor)
		} // if
	} // for

	return errs
}
