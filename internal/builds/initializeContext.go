package builds

import (
	"github.com/yinweli/Sheeter/internal/nameds"
)

// InitializeContext 初始化
func InitializeContext(config *Config) *Context {
	context := &Context{
		Global: &config.Global,
	}

	for _, itor := range config.Elements {
		context.Element = append(context.Element, &initializeElement{
			Global: &config.Global,
			Named:  &nameds.Named{ExcelName: itor.Excel, SheetName: itor.Sheet},
		})
	} // for

	for _, itor := range config.Enums {
		context.Element = append(context.Element, &initializeEnum{
			Global: &config.Global,
			Named:  &nameds.Named{ExcelName: itor.Excel, SheetName: itor.Sheet},
		})
	} // for

	return context
}
