package builds

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// readExcel 讀取excel檔案
func readExcel(content *Content) error {
	excel, err := excelize.OpenFile(content.Excel)

	if err != nil {
		return fmt.Errorf("%s: read excel failed", content.ShowName())
	} // if

	if excel.GetSheetIndex(content.Sheet) == -1 {
		return fmt.Errorf("%s: read excel failed, sheet not found", content.ShowName())
	} // if

	content.excel = excel
	return nil
}
