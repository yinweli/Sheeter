package builder

import (
	"fmt"
	"path"

	"github.com/xuri/excelize/v2"
)

// ReadExcel 讀取excel
func ReadExcel(cargo *Cargo) error {
	file, err := excelize.OpenFile(path.Join(cargo.Global.ExcelPath, cargo.Element.Excel))

	if err != nil {
		return err
	} // if

	defer closeExcel(file)

	rows, err := file.GetRows(cargo.Element.Sheet)

	if err != nil {
		return err
	} // if

	for _, itor := range rows {
		fmt.Println(itor)
	} // for

	return nil
}

func closeExcel(file *excelize.File) {
	err := file.Close()

	if err != nil {
		fmt.Println(err)
	} // if
}
