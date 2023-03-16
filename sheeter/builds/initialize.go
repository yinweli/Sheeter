package builds

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/excels"
	"github.com/yinweli/Sheeter/v2/sheeter/pipelines"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// InitializeData 初始化資料
type InitializeData struct {
	*excels.Excel        // excel物件
	*excels.Sheet        // sheet物件
	ExcelName     string // excel名稱
	SheetName     string // sheet名稱
}

// Initialize 初始化處理
func Initialize(config *Config) (result []*InitializeData, err []error) {
	resultExcel, err := pipelines.Pipeline[string]("search excel", config.Path(), []pipelines.PipelineFunc[string]{
		searchExcel,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	resultSheet, err := pipelines.Pipeline[string]("search sheet", utils.Combine(config.File(), resultExcel), []pipelines.PipelineFunc[string]{
		searchSheet,
	})

	if len(err) > 0 {
		return nil, err
	} // if

	for _, itor := range resultSheet {
		if data, ok := itor.(*InitializeData); ok {
			if utils.CheckExcel(utils.FileName(data.ExcelName)) == false {
				err = append(err, fmt.Errorf("initialize: excel name invalid: %v#%v", data.ExcelName, data.SheetName))
			} // if

			if utils.CheckSheet(data.SheetName) == false {
				err = append(err, fmt.Errorf("initialize: sheet name invalid: %v#%v", data.ExcelName, data.SheetName))
			} // if

			result = append(result, data)
		} // if
	} // for

	if len(err) > 0 {
		return nil, err
	} // if

	return result, nil
}

// searchExcel 搜尋excel
func searchExcel(input string, result chan any) error {
	if err := filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		} // if

		if info.IsDir() {
			return nil
		} // if

		if filepath.Ext(path) != sheeter.ExcelExt {
			return nil
		} // if

		if utils.CheckIgnore(filepath.Base(path)) {
			return nil
		} // if

		result <- path
		return nil
	}); err != nil {
		return fmt.Errorf("search excel: %w", err)
	} // if

	return nil
}

// searchSheet 搜尋sheet
func searchSheet(input string, result chan any) error {
	excel := &excels.Excel{}

	if err := excel.Open(input); err != nil {
		return fmt.Errorf("search sheet: %w", err)
	} // if

	for _, itor := range excel.Sheets() {
		if utils.CheckIgnore(itor) {
			continue
		} // if

		sheet, err := excel.Get(itor)

		if err != nil {
			return fmt.Errorf("search sheet: %w", err)
		} // if

		result <- &InitializeData{
			Excel:     excel,
			Sheet:     sheet,
			ExcelName: filepath.Base(input),
			SheetName: itor,
		}
	} // for

	return nil
}
