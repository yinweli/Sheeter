package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/v2/sheeter/nameds"
	"github.com/yinweli/Sheeter/v2/sheeter/pipelines"
	"github.com/yinweli/Sheeter/v2/sheeter/tmpls"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// PoststepData 後製資料
type PoststepData struct {
	*Config                       // 設定資料
	*nameds.Named                 // 命名工具
	Struct        []*nameds.Named // 結構列表
}

// Poststep 後製處理
func Poststep(config *Config, input []*InitializeData) (file []any, err []error) {
	result := &PoststepData{
		Config: config,
		Named: &nameds.Named{ // 由於只需要AppName與Namespace, 所以不必填寫excel與sheet名稱
			Output: config.Output, // 但是需要表格器路徑, 所以要填寫輸出路徑
		},
	}

	for _, itor := range input {
		result.Struct = append(result.Struct, &nameds.Named{
			ExcelName: itor.ExcelName,
			SheetName: itor.SheetName,
		})
	} // for

	file, err = pipelines.Pipeline[*PoststepData]("poststep", []*PoststepData{result}, []pipelines.PipelineFunc[*PoststepData]{
		generateSheeterCs,
		generateSheeterGo,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	return file, nil
}

// generateSheeterCs 產生cs表格器程式碼
func generateSheeterCs(input *PoststepData, result chan any) error {
	file := input.SheeterPathCs()

	if err := utils.WriteTmpl(file, tmpls.SheeterCs, input); err != nil {
		return fmt.Errorf("generate sheeter cs: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	result <- file
	return nil
}

// generateSheeterGo 產生go表格器程式碼
func generateSheeterGo(input *PoststepData, result chan any) error {
	file := input.SheeterPathGo()

	if err := utils.WriteTmpl(file, tmpls.SheeterGo, input); err != nil {
		return fmt.Errorf("generate sheeter go: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	result <- file
	return nil
}
