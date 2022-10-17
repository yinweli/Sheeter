package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/internal/workflow"
)

// Encoding 產生資料
func Encoding(context *Context) []error {
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

	work := workflow.NewWorkflow("encoding ", totalCount)

	for _, itor := range context.Sector {
		ref := &encodingData{ // 多執行緒需要使用中間變數
			Global:     &context.Config.Global,
			Element:    &itor.Element,
			Mixed:      mixeds.NewMixed(itor.Element.Excel, itor.Element.Sheet),
			excel:      itor.excel,
			layoutJson: itor.layoutJson,
		}

		go func() {
			for _, itor := range tasks {
				if err := itor(ref); err != nil {
					work.Error(fmt.Errorf("encoding failed: %w", err))
				} // if

				work.Increment()
			} // for
		}()
	} // for

	errs := work.End()
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
