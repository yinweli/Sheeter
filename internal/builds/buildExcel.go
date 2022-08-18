package builds

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// readExcel 讀取excel檔案
func readExcel(content *Content) error {
	excel, err := excelize.OpenFile(content.ExcelFilePath())

	if err != nil {
		return fmt.Errorf("%s: read excel failed", content.TargetName())
	} // if

	if excel.GetSheetIndex(content.Sheet) == -1 {
		return fmt.Errorf("%s: read excel failed, sheet not found", content.TargetName())
	} // if

	content.excel = excel
	return nil
}
