package builds

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/names"
	"github.com/yinweli/Sheeter/internal/utils"
)

// SectorInit 區段初始化
func SectorInit(sector *Sector) error {
	sector.named = &names.Named{Excel: sector.Excel, Sheet: sector.Sheet}
	excel, err := excelize.OpenFile(sector.Element.Excel)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, excel can't open: %w", sector.named.StructName(), err)
	} // if

	if excel.GetSheetIndex(sector.Sheet) == -1 {
		return fmt.Errorf("%s: sector init failed, sheet not found", sector.named.StructName())
	} // if

	sector.excel = excel
	fieldLine, err := sector.GetColumns(sector.LineOfField)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, field line not found: %w", sector.named.StructName(), err)
	} // if

	layerLine, err := sector.GetColumns(sector.LineOfLayer)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, layer line not found: %w", sector.named.StructName(), err)
	} // if

	noteLine, err := sector.GetColumns(sector.LineOfNote)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, note line not found: %w", sector.named.StructName(), err)
	} // if

	layoutJson := layouts.NewLayoutJson()
	layoutType := layouts.NewLayoutType()

	if err := layoutType.Begin(sector.named.StructName(), sector.named); err != nil {
		return fmt.Errorf("%s: sector init failed, layoutType begin failed: %w", sector.named.StructName(), err)
	} // if

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("%s: sector init failed, parse field failed: %w", sector.named.StructName(), err)
		} // if

		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: sector init failed, parse layer failed: %w", sector.named.StructName(), err)
		} // if

		note := utils.GetItem(noteLine, col)

		if err := layoutJson.Add(name, field, layer, back); err != nil {
			return fmt.Errorf("%s: sector init failed, layoutJson failed: %w", sector.named.StructName(), err)
		} // if

		if err := layoutType.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: sector init failed, layoutType failed: %w", sector.named.StructName(), err)
		} // if
	} // for

	if err := layoutType.End(); err != nil {
		return fmt.Errorf("%s: sector init failed, layoutType end failed: %w", sector.named.StructName(), err)
	} // if

	pkeyCount := layoutJson.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: sector init failed, pkey duplicate", sector.named.StructName())
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: sector init failed, pkey not found", sector.named.StructName())
	} // if

	sector.layoutJson = layoutJson
	sector.layoutType = layoutType
	return nil
}
