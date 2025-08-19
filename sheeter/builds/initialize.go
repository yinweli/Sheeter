package builds

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/sheeter/pipelines"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Initialize 初始化處理
func Initialize(config *Config) (result []*InitializeData, err []error) {
	resultExcel, err := pipelines.Pipeline[string]("search excel", config.Path(), []pipelines.Execute[string]{
		searchExcel,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	resultSheet, err := pipelines.Pipeline[string]("search sheet", utils.Combine(config.File(), resultExcel), []pipelines.Execute[string]{
		searchSheet,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	for _, itor := range resultSheet {
		data, ok := itor.(*InitializeData)

		if ok == false {
			continue
		} // if

		excel := utils.FileName(data.ExcelName)
		sheet := data.SheetName

		if config.Excluded(excel, sheet) {
			continue
		} // if

		if utils.CheckExcel(excel) == false {
			err = append(err, fmt.Errorf("initialize: excel name invalid: %v#%v", excel, sheet))
		} // if

		if utils.CheckSheet(sheet) == false {
			err = append(err, fmt.Errorf("initialize: sheet name invalid: %v#%v", excel, sheet))
		} // if

		result = append(result, data)
	} // for

	if len(err) > 0 {
		return nil, err
	} // if

	return result, nil
}

// InitializeData 初始化資料
type InitializeData struct {
	*excels.Excel        // excel物件
	*excels.Sheet        // sheet物件
	ExcelName     string // excel名稱
	SheetName     string // sheet名稱
}

// searchExcel 搜尋excel
func searchExcel(material string) (result pipelines.Output) {
	if err := filepath.Walk(material, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		} // if

		if info.IsDir() {
			return nil
		} // if

		if filepath.Ext(path) != sheeter.ExtExcel {
			return nil
		} // if

		if utils.CheckIgnore(filepath.Base(path)) {
			return nil
		} // if

		result.Result = append(result.Result, path)
		return nil
	}); err != nil {
		result.Error = fmt.Errorf("search excel: %w", err)
	} // if

	return result
}

// searchSheet 搜尋sheet
func searchSheet(material string) (result pipelines.Output) {
	excel := &excels.Excel{}

	if err := excel.Open(material); err != nil {
		result.Error = fmt.Errorf("search sheet: %w", err)
		return result
	} // if

	for _, itor := range excel.Sheet() {
		if utils.CheckIgnore(itor) {
			continue
		} // if

		sheet, err := excel.Get(itor)

		if err != nil {
			result.Error = fmt.Errorf("search sheet: %w", err)
			return result
		} // if

		result.Result = append(result.Result, &InitializeData{
			Excel:     excel,
			Sheet:     sheet,
			ExcelName: filepath.Base(material),
			SheetName: itor,
		})
	} // for

	return result
}
