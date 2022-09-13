package builds

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal/fields"
	"github.com/yinweli/Sheeter/internal/layers"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// SectorInit 區段初始化
func SectorInit(sector *Sector) error {
	excel, err := excelize.OpenFile(sector.Excel)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, excel can't open", sector.StructName())
	} // if

	if excel.GetSheetIndex(sector.Sheet) == -1 {
		return fmt.Errorf("%s: sector init failed, sheet not found", sector.StructName())
	} // if

	sector.excel = excel
	fieldLine, err := sector.GetColumns(sector.LineOfField)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, field line not found", sector.StructName())
	} // if

	layerLine, err := sector.GetColumns(sector.LineOfLayer)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, layer line not found", sector.StructName())
	} // if

	/* noteLine TODO: 等待要做layoutType時應該會用到 */
	_, err = sector.GetColumns(sector.LineOfNote)

	if err != nil {
		return fmt.Errorf("%s: sector init failed, note line not found", sector.StructName())
	} // if

	layoutJson := layouts.NewLayoutJson()

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("%s: sector init failed: %w", sector.StructName(), err)
		} // if

		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: sector init failed: %w", sector.StructName(), err)
		} // if

		// note := utils.GetItem(noteLine, col) TODO: 等待要做layoutType時應該會用到

		if err := layoutJson.Add(name, field, layer, back); err != nil {
			return fmt.Errorf("%s: sector init failed: %w", sector.StructName(), err)
		} // if
	} // for

	pkeyCount := layoutJson.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: sector init failed, pkey duplicate", sector.StructName())
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: sector init failed, pkey not found", sector.StructName())
	} // if

	sector.layoutJson = layoutJson

	return nil
}
