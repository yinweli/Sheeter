package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/utils"
)

// initializeSheetEnum 初始化列舉表格資料
type initializeSheetEnum struct {
	*Global                           // 全域設定
	*nameds.Named                     // 命名工具
	excel         *excels.Excel       // excel物件
	layoutEnum    *layouts.LayoutEnum // 列舉布局器
}

// InitializeSheetEnum 初始化列舉表格
func InitializeSheetEnum(material any, _ chan any) error {
	data, ok := material.(*initializeSheetEnum)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if utils.NameCheck(structName) == false {
		return fmt.Errorf("%s: initialize sheetEnum failed: invalid excel & sheet name", structName)
	} // if

	if utils.NameKeywords(structName) == false {
		return fmt.Errorf("%s: initialize sheetEnum failed: conflict with keywords", structName)
	} // if

	if data.excel == nil {
		excel := &excels.Excel{}

		if err := excel.Open(data.ExcelName); err != nil {
			return fmt.Errorf("%s: initialize sheetEnum failed: %w", structName, err)
		} // if

		data.excel = excel
	} // if

	sheet, err := data.excel.Get(data.SheetName)

	if err != nil {
		return fmt.Errorf("%s: initialize sheetEnum failed: sheet not found", structName)
	} // if

	sheet.Nextn(data.LineOfEnum)
	layoutEnum := layouts.NewLayoutEnum()

	for next := true; next; next = sheet.Next() {
		line, _ := sheet.Data()

		if line == nil { // 碰到空行就結束了
			break
		} // if

		if err := layoutEnum.Add(line); err != nil {
			return fmt.Errorf("%s: initialize sheetEnum failed: %w", structName, err)
		} // if
	} // for

	data.layoutEnum = layoutEnum
	return nil
}
