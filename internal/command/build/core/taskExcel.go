package core

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

// executeExcel 讀取excel檔案並獲取表格列表
func (this *Task) executeExcel() error {
	file, err := os.Open(this.excelFilePath())

	if err != nil {
		return fmt.Errorf("file open failed/file not found: %s", this.logName())
	} // if

	defer func() { _ = file.Close() }()
	excel, err := excelize.OpenReader(file)

	if err != nil {
		return fmt.Errorf("excel read failed: %s", this.logName())
	} // if

	this.excel = excel

	if this.sheetExists() == false {
		return fmt.Errorf("sheet not found: %s", this.logName())
	} // if

	return nil
}
