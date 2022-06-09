package core

import (
	"fmt"
	"os"

	"Sheeter/internal/util"

	"github.com/xuri/excelize/v2"
)

// TaskExcel 讀取excel檔案並獲取表格列表
func TaskExcel(ctx *Context) error {
	file, err := os.Open(ctx.ExcelFilePath())

	if err != nil {
		return fmt.Errorf("file open failed/file not found: %s", ctx.LogName())
	} // if

	defer util.SilentClose(file)
	excel, err := excelize.OpenReader(file)

	if err != nil {
		return fmt.Errorf("excel read failed: %s", ctx.LogName())
	} // if

	ctx.Excel = excel

	if ctx.IsSheetExists() == false {
		return fmt.Errorf("sheet not found: %s", ctx.LogName())
	} // if

	return nil
}
