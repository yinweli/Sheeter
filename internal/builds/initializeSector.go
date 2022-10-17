package builds

import (
	"fmt"

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

	if utils.NameCheck(structName) == false {
		return fmt.Errorf("%s: initialize sector failed: invalid excel & sheet name", structName)
	} // if

	if utils.NameKeywords(structName) == false {
		return fmt.Errorf("%s: initialize sector failed: conflict with keywords", structName)
	} // if

	excel := &excels.Excel{}

	if err := excel.Open(sector.Excel); err != nil {
		return fmt.Errorf("%s: initialize sector failed: open excel failed: %w", structName, err)
	} // if

	if excel.Exist(sector.Sheet) == false {
		return fmt.Errorf("%s: initialize sector failed: sheet not found", structName)
	} // if

	line, err := excel.GetLine(
		sector.Sheet,
		context.Global.LineOfName,
		context.Global.LineOfNote,
		context.Global.LineOfField,
		context.Global.LineOfLayer,
	)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed: get line failed: %w", structName, err)
	} // if

	nameLine := line[context.Global.LineOfName]
	noteLine := line[context.Global.LineOfNote]
	fieldLine := line[context.Global.LineOfField]
	layerLine := line[context.Global.LineOfLayer]
	layoutJson := layouts.NewLayoutJson()
	layoutType := layouts.NewLayoutType()
	layoutDepend := layouts.NewLayoutDepend()

	if err := layoutType.Begin(structName, sector.Excel, sector.Sheet); err != nil {
		return fmt.Errorf("%s: initialize sector failed: layoutType begin failed: %w", structName, err)
	} // if

	if err := layoutDepend.Begin(structName); err != nil {
		return fmt.Errorf("%s: initialize sector failed: layoutDepend begin failed: %w", structName, err)
	} // if

	for col, itor := range nameLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name := itor
		note := utils.GetItem(noteLine, col)

		if utils.NameCheck(name) == false {
			return fmt.Errorf("%s: initialize sector failed: invalid name", structName)
		} // if

		field, tag, err := fields.Parser(utils.GetItem(fieldLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize sector failed: parse field failed: %w", structName, err)
		} // if

		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize sector failed: parse layer failed: %w", structName, err)
		} // if

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
