package builds

import (
	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/sheeter/pipelines"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Initialize 初始化處理
func Initialize(config *Config) (context *Context, errs []error) {
	path := preparePath(config.Path())
	excel, errs := pipelines.Pipeline("initialize file", path, []pipelines.PipelineFunc{
		InitializeFile,
	})

	if len(errs) > 0 {
		return nil, errs
	} // if

	excel = prepareExcel(config.Excel(), excel)
	sheet, errs := pipelines.Pipeline("initialize excel", excel, []pipelines.PipelineFunc{
		InitializeExcel,
	})

	if len(errs) > 0 {
		return nil, errs
	} // if

	sheet = prepareSheet(config.Sheet(), sheet, &config.Global) // 由於InitializePick會用到sheet列表, 所以必須在外面準備好列表
	_, errs = pipelines.Pipeline("initialize sheet", sheet, []pipelines.PipelineFunc{
		InitializeSheetData,
		InitializeSheetEnum,
	})

	if len(errs) > 0 {
		return nil, errs
	} // if

	context = &Context{
		Global: &config.Global,
	}

	if err := InitializePick(sheet, context); err != nil {
		return nil, []error{err}
	} // if

	return context, errs
}

// preparePath 準備路徑列表
func preparePath(config []string) []any {
	result := []any{}
	duplicate := utils.NewDuplicate()

	for _, itor := range config {
		if duplicate.Check(itor) {
			result = append(result, itor)
		} // if
	} // for

	return result
}

// prepareExcel 準備excel列表
func prepareExcel(config []string, native []any) []any {
	result := []any{}
	duplicate := utils.NewDuplicate()

	for _, itor := range config {
		if duplicate.Check(itor) {
			result = append(result, itor)
		} // if
	} // for

	for _, itor := range native {
		if value, ok := itor.(string); ok {
			if duplicate.Check(value) {
				result = append(result, value)
			} // if
		} // if
	} // for

	return result
}

// prepareSheet 準備sheet列表
func prepareSheet(config []Sheet, native []any, global *Global) []any {
	result := []any{}
	duplicate := utils.NewDuplicate()

	for _, itor := range config {
		if duplicate.Check(itor.ExcelName, itor.SheetName) {
			if utils.IsDataSheetName(itor.SheetName) {
				result = append(result, &initializeSheetData{
					Named: &nameds.Named{ExcelName: itor.ExcelName, SheetName: itor.SheetName},
				})
			} // if

			if utils.IsEnumSheetName(itor.SheetName) {
				result = append(result, &initializeSheetEnum{
					Named: &nameds.Named{ExcelName: itor.ExcelName, SheetName: itor.SheetName},
				})
			} // if
		} // if
	} // for

	for _, itor := range native {
		if value, ok := itor.(*initializeSheetData); ok {
			if duplicate.Check(value.ExcelName, value.SheetName) {
				result = append(result, itor)
			} // if
		} // if

		if value, ok := itor.(*initializeSheetEnum); ok {
			if duplicate.Check(value.ExcelName, value.SheetName) {
				result = append(result, itor)
			} // if
		} // if
	} // for

	for _, itor := range result {
		if value, ok := itor.(*initializeSheetData); ok {
			value.Global = global
		} // if

		if value, ok := itor.(*initializeSheetEnum); ok {
			value.Global = global
		} // if
	} // for

	return result
}
