package reader

import (
	"fmt"
	"path"

	"Sheeter/internal/command/build/cargo"
	"Sheeter/internal/command/build/field"

	"github.com/xuri/excelize/v2"
)

// ReadSheet 讀取表格
func ReadSheet(cargo *cargo.Cargo) error {
	file, err := excelize.OpenFile(path.Join(cargo.Global.ExcelPath, cargo.Element.Excel))

	if err != nil {
		return err
	} // if

	defer func() {
		_ = file.Close()
	}()

	sheet, err := file.GetRows(cargo.Element.Sheet)

	if sheet == nil || err != nil {
		return fmt.Errorf("sheet not found: %s(%s)", cargo.Element.Excel, cargo.Element.Sheet)
	} // if

	if len(sheet) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return fmt.Errorf("sheet have too less line: %s(%s)", cargo.Element.Excel, cargo.Element.Sheet)
	} // if

	fields := sheet[cargo.Global.GetLineOfField()]

	for _, itor := range fields {
		name, field, err := field.Parse(itor)

		if err != nil {
			return fmt.Errorf("sheet field parse failed: %s(%s), %s", cargo.Element.Excel, cargo.Element.Sheet, err)
		} // if

		fmt.Printf("%s # %s", name, field.TypeExcel())
	} // for

	return nil
}
