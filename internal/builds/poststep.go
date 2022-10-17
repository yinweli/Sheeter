package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/internal/workflow"
)

// Poststep 後製
func Poststep(context *Context) []error {
	tasks := []func(*poststepData) error{}

	if context.Global.ExportJson {
		tasks = append(
			tasks,
			poststepJsonCsDepot,
			poststepJsonGoDepot,
		)
	} // if

	if context.Global.ExportProto {
		tasks = append(
			tasks,
			poststepProtoCsDepot,
			poststepProtoGoDepot,
			poststepProtoCsBat,
			poststepProtoCsSh,
			poststepProtoGoBat,
			poststepProtoGoSh,
		)
	} // if

	totalCount := len(tasks)

	if totalCount <= 0 {
		return []error{}
	} // if

	work := workflow.NewWorkflow("poststep ", totalCount)
	data := &poststepData{
		Global: &context.Config.Global,
		Mixed:  mixeds.NewMixed("", ""),
	}

	for _, itor := range context.Struct {
		data.Struct = append(data.Struct, poststepStruct{
			Mixed: mixeds.NewMixed(itor.types.Excel, itor.types.Sheet),
			Type:  itor.types,
		})
	} // for

	for _, itor := range tasks {
		if err := itor(data); err != nil {
			work.Error(fmt.Errorf("poststep failed: %w", err))
			work.Abort()
			break
		} // if

		work.Increment()
	} // for

	errs := work.End()
	return errs
}

// poststepData 後製資料
type poststepData struct {
	*Global                        // 全域設定
	*mixeds.Mixed                  // 綜合工具
	Struct        []poststepStruct // 結構列表
}

// poststepStruct 後製結構資料
type poststepStruct struct {
	*mixeds.Mixed // 結構列表
	*layouts.Type // 類型資料
}
