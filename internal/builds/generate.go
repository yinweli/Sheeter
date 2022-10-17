package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/internal/workflow"
)

// Generate 產生程式碼
func Generate(context *Context) []error {
	tasks := []func(*generateData) error{}

	if context.Global.ExportJson {
		tasks = append(
			tasks,
			generateJsonStructCs,
			generateJsonReaderCs,
			generateJsonStructGo,
			generateJsonReaderGo,
		)
	} // if

	if context.Global.ExportProto {
		tasks = append(
			tasks,
			generateProtoSchema,
			generateProtoReaderCs,
			generateProtoReaderGo,
		)
	} // if

	itemCount := len(context.Struct)
	taskCount := len(tasks)
	totalCount := itemCount * taskCount

	if totalCount <= 0 {
		return []error{}
	} // if

	work := workflow.NewWorkflow("generate ", totalCount)

	for _, itor := range context.Struct {
		ref := &generateData{ // 多執行緒需要使用中間變數
			Global: &context.Config.Global,
			Mixed:  mixeds.NewMixed(itor.types.Excel, itor.types.Sheet),
			Type:   itor.types,
			Depend: itor.depend,
		}

		go func() {
			for _, itor := range tasks {
				if err := itor(ref); err != nil {
					work.Error(fmt.Errorf("generate failed: %w", err))
				} // if

				work.Increment()
			} // for
		}()
	} // for

	errs := work.End()
	return errs
}

// generateData 產生程式碼資料
type generateData struct {
	*Global                // 全域設定
	*mixeds.Mixed          // 綜合工具
	*layouts.Type          // 類型資料
	Depend        []string // 依賴列表
}
