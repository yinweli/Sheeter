package builds

import (
	"fmt"

	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
)

// Poststep 後製
func Poststep(context *Context) error {
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
	progress := mpb.New(mpb.WithWidth(internal.BarWidth))
	progressbar := progress.AddBar(
		int64(totalCount),
		mpb.PrependDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Name("poststep "),
			decor.OnComplete(decor.Spinner(nil), "complete"),
		),
	)

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
			return fmt.Errorf("poststep failed: %w", err)
		} // if

		progressbar.Increment()
	} // for

	progress.Wait()
	return nil
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
