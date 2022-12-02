package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/sheeter/fields"
	"github.com/yinweli/Sheeter/sheeter/layers"
	"github.com/yinweli/Sheeter/sheeter/layouts"
	"github.com/yinweli/Sheeter/sheeter/nameds"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// initializeSheetData 初始化資料表格資料
type initializeSheetData struct {
	*Global                             // 全域設定
	*nameds.Named                       // 命名工具
	excel         *excels.Excel         // excel物件
	layoutData    *layouts.LayoutData   // 資料布局器
	layoutType    *layouts.LayoutType   // 類型布局器
	layoutDepend  *layouts.LayoutDepend // 依賴布局器
}

// InitializeSheetData 初始化資料表格
func InitializeSheetData(material any, _ chan any) error {
	data, ok := material.(*initializeSheetData)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if utils.NameCheck(structName) == false {
		return fmt.Errorf("%s: initialize sheetData failed: invalid excel & sheet name", structName)
	} // if

	if utils.NameKeywords(structName) == false {
		return fmt.Errorf("%s: initialize sheetData failed: conflict with keywords", structName)
	} // if

	if data.excel == nil {
		excel := &excels.Excel{}

		if err := excel.Open(data.ExcelName); err != nil {
			return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
		} // if

		data.excel = excel
	} // if

	line, err := data.excel.GetLine(
		data.SheetName,
		data.LineOfTag,
		data.LineOfName,
		data.LineOfNote,
		data.LineOfField,
		data.LineOfLayer,
	)

	if err != nil {
		return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
	} // if

	tagLine := line[data.LineOfTag]
	nameLine := line[data.LineOfName] // 尋訪時, 以名稱行為主
	noteLine := line[data.LineOfNote]
	fieldLine := line[data.LineOfField]
	layerLine := line[data.LineOfLayer]
	layoutData := layouts.NewLayoutData()
	layoutType := layouts.NewLayoutType()
	layoutDepend := layouts.NewLayoutDepend()

	if err := layoutType.Begin(structName, data.ExcelName, data.SheetName); err != nil {
		return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
	} // if

	if err := layoutDepend.Begin(structName); err != nil {
		return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
	} // if

	for col, itor := range nameLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name := itor
		note := utils.GetItem(noteLine, col)

		if utils.NameCheck(name) == false {
			return fmt.Errorf("%s: initialize sheetData failed: invalid name", structName)
		} // if

		field, err := fields.Parser(utils.GetItem(fieldLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
		} // if

		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
		} // if

		tag := utils.GetItem(tagLine, col)

		if err := layoutData.Add(name, field, layer, back, tag); err != nil {
			return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
		} // if

		if err := layoutType.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
		} // if

		if err := layoutDepend.Add(layer, back); err != nil {
			return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
		} // if
	} // for

	if err := layoutType.End(); err != nil {
		return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
	} // if

	if err := layoutDepend.End(); err != nil {
		return fmt.Errorf("%s: initialize sheetData failed: %w", structName, err)
	} // if

	pkeyCount := layoutData.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: initialize sheetData failed: pkey duplicate", structName)
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: initialize sheetData failed: pkey not found", structName)
	} // if

	data.layoutData = layoutData
	data.layoutType = layoutType
	data.layoutDepend = layoutDepend
	return nil
}
