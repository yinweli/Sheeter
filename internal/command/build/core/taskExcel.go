package core

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// executeExcel 讀取excel檔案並獲取表格列表
func (this *Task) executeExcel() error {
	excel, err := excelize.OpenFile(this.excelFilePath())

	if err != nil {
		return fmt.Errorf("excel read failed: %s", this.originalName())
	} // if

	this.excel = excel

	if this.sheetExists() == false {
		return fmt.Errorf("sheet not found: %s", this.originalName())
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressM)
	} // if

	return nil
}
