package builds

import (
	"fmt"
	"sort"

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
	material := &PoststepData{
		Config: config,
		Named: &nameds.Named{ // 由於只需要AppName與Namespace, 所以不必填寫excel與sheet名稱
			Output: config.Output, // 但是需要表格器路徑, 所以要填寫輸出路徑
		},
	}

	for _, itor := range input {
		material.Struct = append(material.Struct, &nameds.Named{
			ExcelName: itor.ExcelName,
			SheetName: itor.SheetName,
		})
	} // for

	sort.Slice(material.Struct, func(l, r int) bool { // 經過排序後讓產生程式碼時能夠更加一致
		lhs := material.Struct[l]
		rhs := material.Struct[r]

		if lhs.ExcelName == rhs.ExcelName {
			return lhs.SheetName < rhs.SheetName
		} // if

		return lhs.ExcelName < rhs.ExcelName
	})

	file, err = pipelines.Pipeline[*PoststepData]("poststep", []*PoststepData{material}, []pipelines.PipelineFunc[*PoststepData]{
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
	path := input.SheeterPathCs()

	if err := utils.WriteTmpl(path, tmpls.SheeterCs, input); err != nil {
		return fmt.Errorf("generate sheeter cs: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	result <- path
	return nil
}

// generateSheeterGo 產生go表格器程式碼
func generateSheeterGo(input *PoststepData, result chan any) error {
	path := input.SheeterPathGo()

	if err := utils.WriteTmpl(path, tmpls.SheeterGo, input); err != nil {
		return fmt.Errorf("generate sheeter go: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	result <- path
	return nil
}
