package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/v2/sheeter/excels"
	"github.com/yinweli/Sheeter/v2/sheeter/layouts"
	"github.com/yinweli/Sheeter/v2/sheeter/nameds"
	"github.com/yinweli/Sheeter/v2/sheeter/pipelines"
	"github.com/yinweli/Sheeter/v2/sheeter/tmpls"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// OperationData 作業資料
type OperationData struct {
	*Global                         // 全域設定
	*excels.Excel                   // excel物件
	*excels.Sheet                   // sheet物件
	*nameds.Named                   // 命名工具
	*nameds.Pkey                    // 主要索引命名工具
	Field           []*nameds.Field // 欄位列表
	*layouts.Layout                 // 布局資料
}

// Operation 作業處理
func Operation(config *Config, input []*InitializeData) (file []any, err []error) {
	result := []*OperationData{}

	for _, itor := range input {
		result = append(result, &OperationData{
			Global: &config.Global,
			Excel:  itor.Excel,
			Sheet:  itor.Sheet,
			Named: &nameds.Named{
				ExcelName: itor.ExcelName,
				SheetName: itor.SheetName,
			},
		})
	} // for

	file, err = pipelines.Pipeline[*OperationData]("operation", result, []pipelines.PipelineFunc[*OperationData]{
		parseLayout,
		generateData,
		generateReaderCs,
		generateReaderGo,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	return file, nil
}

// parseLayout 解析布局
func parseLayout(input *OperationData, _ chan any) error {
	line, err := input.GetLine(
		input.SheetName,
		input.LineOfTag,
		input.LineOfName,
		input.LineOfNote,
		input.LineOfField,
	)

	if err != nil {
		return fmt.Errorf("parse layout: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	layout := layouts.NewLayout(input.AutoKey)
	lineTag := line[input.LineOfTag]
	lineName := line[input.LineOfName]
	lineNote := line[input.LineOfNote]
	lineField := line[input.LineOfField]

	if err = layout.Set(lineTag, lineName, lineNote, lineField); err != nil {
		return fmt.Errorf("parse layout: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	pkey := layout.Pkey(input.Tag)

	if pkey == nil {
		return fmt.Errorf("parse layout: %v#%v: pkey not exist", input.ExcelName, input.SheetName)
	} // if

	input.Pkey = &nameds.Pkey{
		Pkey: pkey.Field,
	}

	for _, itor := range layout.Layout(input.Tag) {
		input.Field = append(input.Field, &nameds.Field{
			Data: itor,
		})
	} // for

	input.Layout = layout
	return nil
}

// generateData 產生資料檔案
func generateData(input *OperationData, result chan any) error {
	json, err := layouts.JsonPack(input.Tag, input.LineOfData, input.Sheet, input.Layout)

	if err != nil {
		return fmt.Errorf("generate data: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	file := input.DataPath()

	if err := utils.WriteFile(file, json); err != nil {
		return fmt.Errorf("generate data: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	result <- file
	return nil
}

// generateReaderCs 產生cs讀取器程式碼
func generateReaderCs(input *OperationData, result chan any) error {
	file := input.ReaderPathCs()

	if err := utils.WriteTmpl(file, tmpls.ReaderCs, input); err != nil {
		return fmt.Errorf("generate reader cs: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	result <- file
	return nil
}

// generateReaderGo 產生go讀取器程式碼
func generateReaderGo(input *OperationData, result chan any) error {
	file := input.ReaderPathGo()

	if err := utils.WriteTmpl(file, tmpls.ReaderGo, input); err != nil {
		return fmt.Errorf("generate reader go: %v#%v: %w", input.ExcelName, input.SheetName, err)
	} // if

	result <- file
	return nil
}
