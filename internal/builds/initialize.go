package builds

import (
	"strings"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/pipelines"
)

// Initialize 初始化處理
func Initialize(config *Config) (context *Context, errs []error) {
	path := []any{}

	for _, itor := range config.Path() {
		path = append(path, itor)
	} // for

	excel, errs := pipelines.Pipeline("initialize file", path, []pipelines.PipelineFunc{
		InitializeFile,
	})

	if len(errs) > 0 {
		return nil, errs
	} // if

	for _, itor := range config.Excel() {
		excel = append(excel, itor)
	} // for

	sheet, errs := pipelines.Pipeline("initialize excel", excel, []pipelines.PipelineFunc{
		InitializeExcel,
	})

	if len(errs) > 0 {
		return nil, errs
	} // if

	for _, itor := range config.Sheet() {
		if strings.HasPrefix(itor.SheetName, internal.SignData) {
			sheet = append(sheet, &initializeSheetData{
				Named: &nameds.Named{ExcelName: itor.ExcelName, SheetName: itor.SheetName},
			})
		} // if

		if strings.HasPrefix(itor.SheetName, internal.SignEnum) {
			sheet = append(sheet, &initializeSheetEnum{
				Named: &nameds.Named{ExcelName: itor.ExcelName, SheetName: itor.SheetName},
			})
		} // if
	} // for

	for _, itor := range sheet {
		if value, ok := itor.(*initializeSheetData); ok {
			value.Global = &config.Global
		} // if

		if value, ok := itor.(*initializeSheetEnum); ok {
			value.Global = &config.Global
		} // if
	} // for

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
