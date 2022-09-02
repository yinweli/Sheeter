package builds

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal/builds/fields"
	"github.com/yinweli/Sheeter/internal/builds/layers"
	"github.com/yinweli/Sheeter/internal/builds/layouts"
	"github.com/yinweli/Sheeter/internal/util"
)

// Initialize 初始化
func Initialize(content *Content) error {
	excel, err := excelize.OpenFile(content.Excel)

	if err != nil {
		return fmt.Errorf("%s: initialize failed, excel can't open", content.ShowName())
	} // if

	if excel.GetSheetIndex(content.Sheet) == -1 {
		return fmt.Errorf("%s: initialize failed, sheet not found", content.ShowName())
	} // if

	content.excel = excel
	fieldLine, err := content.GetColumns(content.LineOfField)

	if err != nil {
		return fmt.Errorf("%s: initialize failed, field line not found", content.ShowName())
	} // if

	layerLine, err := content.GetColumns(content.LineOfLayer)

	if err != nil {
		return fmt.Errorf("%s: initialize failed, layer line not found", content.ShowName())
	} // if

	noteLine, err := content.GetColumns(content.LineOfNote)

	if err != nil {
		return fmt.Errorf("%s: initialize failed, note line not found", content.ShowName())
	} // if

	builder := layouts.NewBuilder()

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("%s: initialize failed: %w", content.ShowName(), err)
		} // if

		layer, back, err := layers.Parser(util.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize failed: %w", content.ShowName(), err)
		} // if

		note := util.GetItem(noteLine, col)

		if err := builder.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: initialize failed: %w", content.ShowName(), err)
		} // if
	} // for

	pkeyCount := builder.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: initialize failed, pkey duplicate", content.ShowName())
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: initialize failed, pkey not found", content.ShowName())
	} // if

	content.builder = builder
	return nil
}
