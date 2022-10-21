package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/nameds"
	"github.com/yinweli/Sheeter/internal/utils"
)

// initializeElement 初始化項目資料
type initializeElement struct {
	*Global                             // 全域設定
	*nameds.Named                       // 命名工具
	excel         *excels.Excel         // excel物件
	layoutData    *layouts.LayoutData   // 資料布局器
	layoutType    *layouts.LayoutType   // 類型布局器
	layoutDepend  *layouts.LayoutDepend // 依賴布局器
}

// Close 關閉初始化項目資料
func (this *initializeElement) Close() {
	if this.excel != nil {
		this.excel.Close()
	} // if
}

// InitializeElement 初始化項目
func InitializeElement(material any) error {
	data, ok := material.(*initializeElement)

	if ok == false {
		return nil
	} // if

	structName := data.StructName()

	if utils.NameCheck(structName) == false {
		return fmt.Errorf("%s: initialize element failed: invalid excel & sheet name", structName)
	} // if

	if utils.NameKeywords(structName) == false {
		return fmt.Errorf("%s: initialize element failed: conflict with keywords", structName)
	} // if

	excel := &excels.Excel{}

	if err := excel.Open(data.ExcelName); err != nil {
		return fmt.Errorf("%s: initialize element failed: open excel failed: %w", structName, err)
	} // if

	if excel.Exist(data.SheetName) == false {
		return fmt.Errorf("%s: initialize element failed: sheet not found", structName)
	} // if

	line, err := excel.GetLine(
		data.SheetName,
		data.LineOfName,
		data.LineOfNote,
		data.LineOfField,
		data.LineOfLayer,
	)

	if err != nil {
		return fmt.Errorf("%s: initialize element failed: get line failed: %w", structName, err)
	} // if

	nameLine := line[data.LineOfName]
	noteLine := line[data.LineOfNote]
	fieldLine := line[data.LineOfField]
	layerLine := line[data.LineOfLayer]
	layoutData := layouts.NewLayoutData()
	layoutType := layouts.NewLayoutType()
	layoutDepend := layouts.NewLayoutDepend()

	if err := layoutType.Begin(structName, data.ExcelName, data.SheetName); err != nil {
		return fmt.Errorf("%s: initialize element failed: layoutType begin failed: %w", structName, err)
	} // if

	if err := layoutDepend.Begin(structName); err != nil {
		return fmt.Errorf("%s: initialize element failed: layoutDepend begin failed: %w", structName, err)
	} // if

	for col, itor := range nameLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name := itor
		note := utils.GetItem(noteLine, col)

		if utils.NameCheck(name) == false {
			return fmt.Errorf("%s: initialize element failed: invalid name", structName)
		} // if

		field, tag, err := fields.Parser(utils.GetItem(fieldLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize element failed: parse field failed: %w", structName, err)
		} // if

		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize element failed: parse layer failed: %w", structName, err)
		} // if

		if err := layoutData.Add(name, field, tag, layer, back); err != nil {
			return fmt.Errorf("%s: initialize element failed: layoutData add failed: %w", structName, err)
		} // if

		if err := layoutType.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: initialize element failed: layoutType add failed: %w", structName, err)
		} // if

		if err := layoutDepend.Add(layer, back); err != nil {
			return fmt.Errorf("%s: initialize element failed: layoutDepend add failed: %w", structName, err)
		} // if
	} // for

	if err := layoutType.End(); err != nil {
		return fmt.Errorf("%s: initialize element failed: layoutType end failed: %w", structName, err)
	} // if

	if err := layoutDepend.End(); err != nil {
		return fmt.Errorf("%s: initialize element failed: layoutDepend end failed: %w", structName, err)
	} // if

	pkeyCount := layoutData.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: initialize element failed: pkey duplicate", structName)
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: initialize element failed: pkey not found", structName)
	} // if

	data.excel = excel
	data.layoutData = layoutData
	data.layoutType = layoutType
	data.layoutDepend = layoutDepend
	return nil
}
