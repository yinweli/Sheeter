package builds

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/nameds"
)

// InitializeExcel 初始化excel
func InitializeExcel(material any, result chan any) error {
	data, ok := material.(string)

	if ok == false {
		return nil
	} // if

	excel := &excels.Excel{}

	if err := excel.Open(data); err != nil {
		return fmt.Errorf("%s: initialize excel failed: %w", data, err)
	} // if

	for _, itor := range excel.Sheets() {
		if strings.HasPrefix(itor, internal.SignData) {
			result <- &initializeSheetData{
				Named: &nameds.Named{ExcelName: data, SheetName: itor},
				excel: excel,
			}
		} // if

		if strings.HasPrefix(itor, internal.SignEnum) {
			result <- &initializeSheetEnum{
				Named: &nameds.Named{ExcelName: data, SheetName: itor},
				excel: excel,
			}
		} // if
	} // for

	return nil
}
