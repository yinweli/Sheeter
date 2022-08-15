package tasks

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// excel 讀取excel檔案並獲取表格列表
func (this *Task) excel() error {
	excel, err := excelize.OpenFile(this.excelFilePath())

	if err != nil {
		return fmt.Errorf("excel failed: %s", this.targetName())
	} // if

	this.xlsfile = excel

	if this.sheetExists() == false {
		return fmt.Errorf("excel failed: %s\nsheet not found", this.targetName())
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
