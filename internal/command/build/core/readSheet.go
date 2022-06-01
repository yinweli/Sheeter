package core

import (
	"fmt"
	"io"
	"path"

	"github.com/xuri/excelize/v2"
)

// ReadSheet 讀取表格
func ReadSheet(cargo *Cargo, writer io.Writer) error {
	file, err := excelize.OpenFile(path.Join(cargo.Global.ExcelPath, cargo.Element.Excel))

	if err != nil {
		return err
	} // if

	defer func() {
		_ = file.Close()
	}()

	sheets, err := file.GetRows(cargo.Element.Sheet)

	if sheets == nil || err != nil {
		return fmt.Errorf("sheet not found: %s", cargo.LogName())
	} // if

	if len(sheets) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return fmt.Errorf("sheet have too less line: %s", cargo.LogName())
	} // if

	cargo.Progress = NewProgress(99, cargo.LogName(), writer) // TODO: 計算進度值
	cargo.Sheets = sheets

	return nil
}
