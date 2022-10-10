package builds

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/internal/utils"
)

// initializeSector 初始化區段
func initializeSector(context *Context, sector *ContextSector) error {
	structName := mixeds.NewMixed(sector.Excel, sector.Sheet).StructName()

	for _, itor := range internal.Keywords {
		if strings.EqualFold(structName, itor) {
			return fmt.Errorf("%s: initialize sector failed: invalid excel & sheet name", structName)
		} // if
	} // for

	excel := &excels.Excel{}

	if err := excel.Open(sector.Excel); err != nil {
		return fmt.Errorf("%s: initialize sector failed: open excel failed: %w", structName, err)
	} // if

	if excel.SheetExist(sector.Sheet) == false {
		return fmt.Errorf("%s: initialize sector failed: open excel failed: sheet not found", structName)
	} // if

	fieldLine, err := excel.GetData(sector.Sheet, context.Global.LineOfField)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed: field line not found: %w", structName, err)
	} // if

	layerLine, err := excel.GetData(sector.Sheet, context.Global.LineOfLayer)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed: layer line not found: %w", structName, err)
	} // if

	noteLine, err := excel.GetData(sector.Sheet, context.Global.LineOfNote)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed: note line not found: %w", structName, err)
	} // if

	layoutJson := layouts.NewLayoutJson()
	layoutType := layouts.NewLayoutType()
	layoutDepend := layouts.NewLayoutDepend()

	if err := layoutType.Begin(structName, sector.Excel, sector.Sheet); err != nil {
		return fmt.Errorf("%s: initialize sector failed: layoutType begin failed: %w", structName, err)
	} // if

	if err := layoutDepend.Begin(structName); err != nil {
		return fmt.Errorf("%s: initialize sector failed: layoutDepend begin failed: %w", structName, err)
	} // if

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, tag, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("%s: initialize sector failed: parse field failed: %w", structName, err)
		} // if

		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize sector failed: parse layer failed: %w", structName, err)
		} // if

		note := utils.GetItem(noteLine, col)

		if err := layoutJson.Add(name, field, tag, layer, back); err != nil {
			return fmt.Errorf("%s: initialize sector failed: layoutJson add failed: %w", structName, err)
		} // if

		if err := layoutType.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: initialize sector failed: layoutType add failed: %w", structName, err)
		} // if

		if err := layoutDepend.Add(layer, back); err != nil {
			return fmt.Errorf("%s: initialize sector failed: layoutDepend add failed: %w", structName, err)
		} // if
	} // for

	if err := layoutType.End(); err != nil {
		return fmt.Errorf("%s: initialize sector failed: layoutType end failed: %w", structName, err)
	} // if

	if err := layoutDepend.End(); err != nil {
		return fmt.Errorf("%s: initialize sector failed: layoutDepend end failed: %w", structName, err)
	} // if

	pkeyCount := layoutJson.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: initialize sector failed: pkey duplicate", structName)
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: initialize sector failed: pkey not found", structName)
	} // if

	sector.excel = excel
	sector.layoutJson = layoutJson
	sector.layoutType = layoutType
	sector.layoutDepend = layoutDepend
	return nil
}
