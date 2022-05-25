package core

import (
	"fmt"
	"path"

	"github.com/xuri/excelize/v2"
)

// ReadSheet 讀取表格
func ReadSheet(task *Task) error {
	file, err := excelize.OpenFile(path.Join(task.Global.ExcelPath, task.Element.Excel))

	if err != nil {
		return err
	} // if

	defer func() {
		_ = file.Close()
	}()

	sheet, err := file.GetRows(task.Element.Sheet)

	if sheet == nil || err != nil {
		return fmt.Errorf("sheet not found: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	if len(sheet) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return fmt.Errorf("sheet have too less line: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	fields := sheet[task.Global.GetLineOfField()]

	for _, itor := range fields {
		name, field, err := ParseField(itor)

		if err != nil {
			return fmt.Errorf("sheet field parse failed: %s(%s), %s", task.Element.Excel, task.Element.Sheet, err)
		} // if

		fmt.Printf("%s # %s", name, field.TypeExcel())
	} // for

	return nil
}
