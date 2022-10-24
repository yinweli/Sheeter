package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/utils"
)

// initializeEnum 初始化列舉資料
type initializeEnum struct {
	*Global                           // 全域設定
	*nameds.Named                     // 命名工具
	layoutEnum    *layouts.LayoutEnum // 列舉布局器
}

// InitializeEnum 初始化列舉
func InitializeEnum(material any) error {
	data, ok := material.(*initializeEnum)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if utils.NameCheck(structName) == false {
		return fmt.Errorf("%s: initialize enum failed: invalid excel & sheet name", structName)
	} // if

	if utils.NameKeywords(structName) == false {
		return fmt.Errorf("%s: initialize enum failed: conflict with keywords", structName)
	} // if

	excel := &excels.Excel{}

	if err := excel.Open(data.ExcelName); err != nil {
		return fmt.Errorf("%s: initialize enum failed: open excel failed: %w", structName, err)
	} // if

	sheet, err := excel.Get(data.SheetName)

	if err != nil {
		return fmt.Errorf("%s: initialize enum failed: sheet not found", structName)
	} // if

	sheet.Nextn(data.LineOfEnum)
	layoutEnum := layouts.NewLayoutEnum()

	for ok := true; ok; ok = sheet.Next() {
		data, _ := sheet.Data()

		if data == nil { // 碰到空行就結束了
			break
		} // if

		if err := layoutEnum.Add(data); err != nil {
			return fmt.Errorf("%s: initialize enum failed: layoutEnum add failed: %w", structName, err)
		} // if
	} // for

	data.layoutEnum = layoutEnum
	return nil
}
