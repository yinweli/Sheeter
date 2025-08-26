package builds

import (
	"fmt"
	"sort"

	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/sheeter/layouts"
	"github.com/yinweli/Sheeter/v3/sheeter/nameds"
	"github.com/yinweli/Sheeter/v3/sheeter/pipelines"
	"github.com/yinweli/Sheeter/v3/sheeter/tmpls"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Operation 作業處理
func Operation(config *Config, initializeData []*InitializeData) (result []any, err []error) {
	material := []*OperationData{}

	for _, itor := range initializeData {
		material = append(material, &OperationData{
			Config: config,
			Excel:  itor.Excel,
			Sheet:  itor.Sheet,
			Named: &nameds.Named{
				Output:    config.Output,
				ExcelName: itor.ExcelName,
				SheetName: itor.SheetName,
			},
		})
	} // for

	result, err = pipelines.Pipeline[*OperationData]("operation", material, []pipelines.Execute[*OperationData]{
		parseLayout,
		generateData,
		generateReaderCs,
		generateReaderGo,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	return result, nil
}

// OperationData 作業資料
type OperationData struct {
	*Config                         // 設定資料
	*excels.Excel                   // excel物件
	*excels.Sheet                   // sheet物件
	*nameds.Named                   // 命名工具
	Field           []*nameds.Field // 欄位列表
	*layouts.Layout                 // 布局資料
}

// parseLayout 解析布局
func parseLayout(material *OperationData) (result pipelines.Output) {
	line, err := material.GetLine(
		material.SheetName,
		material.LineOfTag,
		material.LineOfName,
		material.LineOfNote,
		material.LineOfField,
	)

	if err != nil {
		result.Error = fmt.Errorf("parse layout: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	lineTag := line[material.LineOfTag]
	lineName := line[material.LineOfName]
	lineNote := line[material.LineOfNote]
	lineField := line[material.LineOfField]
	layout, failed := layouts.NewLayout(lineTag, lineName, lineNote, lineField)

	if len(failed) > 0 {
		result.Error = fmt.Errorf("parse layout: %v#%v: %v", material.ExcelName, material.SheetName, failed)
		return result
	} // if

	material.Named.Primary = layout.Primary()

	if material.Named.Primary == nil {
		result.Error = fmt.Errorf("parse layout: %v#%v: primary not exist", material.ExcelName, material.SheetName)
		return result
	} // if

	for _, itor := range layout.Select(material.Tag) {
		material.Field = append(material.Field, &nameds.Field{
			Layout: itor,
		})
	} // for

	sort.Slice(material.Field, func(l, r int) bool { // 經過排序後讓產生程式碼時能夠更加一致
		lhs := material.Field[l]
		rhs := material.Field[r]
		return lhs.FieldName() < rhs.FieldName()
	})

	material.Layout = layout
	return result
}

// generateData 產生資料檔案
func generateData(material *OperationData) (result pipelines.Output) {
	json, err := layouts.JsonPack(material.Tag, material.LineOfData, material.Sheet, material.Layout)

	if err != nil {
		result.Error = fmt.Errorf("generate data: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	path := material.DataPath()

	if err = utils.WriteFile(path, json); err != nil {
		result.Error = fmt.Errorf("generate data: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	result.Result = append(result.Result, path)
	return result
}

// generateReaderCs 產生cs語言讀取程式碼
func generateReaderCs(material *OperationData) (result pipelines.Output) {
	path := material.ReaderPathCs()

	if err := utils.WriteTmpl(path, tmpls.ReaderCs, material); err != nil {
		result.Error = fmt.Errorf("generate reader cs: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	result.Result = append(result.Result, path)
	return result
}

// generateReaderGo 產生go語言讀取程式碼
func generateReaderGo(material *OperationData) (result pipelines.Output) {
	path := material.ReaderPathGo()

	if err := utils.WriteTmpl(path, tmpls.ReaderGo, material); err != nil {
		result.Error = fmt.Errorf("generate reader go: %v#%v: %w", material.ExcelName, material.SheetName, err)
		return result
	} // if

	result.Result = append(result.Result, path)
	return result
}
