package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/sheeter/layouts"
	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// encodingJson 編碼json資料
type encodingJson struct {
	*Global                           // 全域設定
	*nameds.Named                     // 命名工具
	*nameds.Json                      // json命名工具
	excel         *excels.Excel       // excel物件
	layoutData    *layouts.LayoutData // 資料布局器
}

// EncodingJson 編碼json資料
func EncodingJson(material any, _ chan any) error {
	data, ok := material.(*encodingJson)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()
	sheet, err := data.excel.Get(data.Named.SheetName)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: sheet not found", structName)
	} // if

	sheet.Nextn(data.LineOfData)
	json, err := layouts.JsonPack(sheet, data.layoutData, data.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(data.JsonDataPath(), json); err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	return nil
}