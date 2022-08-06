package tasks

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// runExcel 讀取excel檔案並獲取表格列表
func (this *Task) runExcel() error {
	excel, err := excelize.OpenFile(this.excelFilePath())

	if err != nil {
		return fmt.Errorf("read excel failed: %s", this.originalName())
	} // if

	this.xlsfile = excel

	if this.sheetExists() == false {
		return fmt.Errorf("read excel failed: %s\nsheet not found", this.originalName())
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
