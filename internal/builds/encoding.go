package builds

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Encoding 產生資料
func Encoding(context *Context) (errs []error) {
	tasks := []func(*encodingData) error{}

	if context.Global.ExportJson {
		tasks = append(
			tasks,
			encodingJson,
		)
	} // if

	if context.Global.ExportProto {
		tasks = append(
			tasks,
			encodingProto,
		)
	} // if

	itemCount := len(context.Sector)
	taskCount := len(tasks)
	totalCount := itemCount * taskCount

	if totalCount <= 0 {
		return []error{}
	} // if

	errors := make(chan error, itemCount) // 結果通訊通道, 拿來緩存執行結果(或是錯誤), 最後全部完成後才印出來
	signaler := utils.NewWaitGroup(itemCount)
	progress := mpb.New(mpb.WithWidth(internal.BarWidth), mpb.WithWaitGroup(signaler))
	progressbar := progress.AddBar(
		int64(totalCount),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("encoding "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

	for _, itor := range context.Sector {
		data := &encodingData{ // 多執行緒需要使用中間變數
			Global:     &context.Config.Global,
			Element:    &itor.Element,
			Mixed:      mixeds.NewMixed(itor.Element.Excel, itor.Element.Sheet),
			excel:      itor.excel,
			layoutJson: itor.layoutJson,
		}

		go func() {
			defer signaler.Done()

			for _, itor := range tasks {
				if err := itor(data); err != nil {
					errors <- fmt.Errorf("encoding failed: %w", err)
				} // if

				progressbar.Increment()
			} // for
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

// encodingData 產生資料資料
type encodingData struct {
	*Global                           // 全域設定
	*Element                          // 項目設定
	*mixeds.Mixed                     // 綜合工具
	excel         *excels.Excel       // excel物件
	layoutJson    *layouts.LayoutJson // json布局器
}
