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

// initializeSector 初始化區段
func initializeSector(runtimeSector *RuntimeSector) error {
	runtimeSector.named = &names.Named{Excel: runtimeSector.Excel, Sheet: runtimeSector.Sheet}
	excel, err := excelize.OpenFile(runtimeSector.Element.Excel)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed, excel can't open: %w", runtimeSector.named.StructName(), err)
	} // if

	if excel.GetSheetIndex(runtimeSector.Sheet) == -1 {
		return fmt.Errorf("%s: initialize sector failed, sheet not found", runtimeSector.named.StructName())
	} // if

	runtimeSector.excel = excel
	fieldLine, err := runtimeSector.GetColumns(runtimeSector.LineOfField)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed, field line not found: %w", runtimeSector.named.StructName(), err)
	} // if

	layerLine, err := runtimeSector.GetColumns(runtimeSector.LineOfLayer)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed, layer line not found: %w", runtimeSector.named.StructName(), err)
	} // if

	noteLine, err := runtimeSector.GetColumns(runtimeSector.LineOfNote)

	if err != nil {
		return fmt.Errorf("%s: initialize sector failed, note line not found: %w", runtimeSector.named.StructName(), err)
	} // if

	layoutJson := layouts.NewLayoutJson()
	layoutType := layouts.NewLayoutType()

	if err := layoutType.Begin(runtimeSector.named.StructName(), runtimeSector.named); err != nil {
		return fmt.Errorf("%s: initialize sector failed, layoutType begin failed: %w", runtimeSector.named.StructName(), err)
	} // if

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("%s: initialize sector failed, parse field failed: %w", runtimeSector.named.StructName(), err)
		} // if

		layer, back, err := layers.Parser(utils.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: initialize sector failed, parse layer failed: %w", runtimeSector.named.StructName(), err)
		} // if

		note := utils.GetItem(noteLine, col)

		if err := layoutJson.Add(name, field, layer, back); err != nil {
			return fmt.Errorf("%s: initialize sector failed, layoutJson failed: %w", runtimeSector.named.StructName(), err)
		} // if

		if err := layoutType.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: initialize sector failed, layoutType failed: %w", runtimeSector.named.StructName(), err)
		} // if
	} // for

	if err := layoutType.End(); err != nil {
		return fmt.Errorf("%s: initialize sector failed, layoutType end failed: %w", runtimeSector.named.StructName(), err)
	} // if

	pkeyCount := layoutJson.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: initialize sector failed, pkey duplicate", runtimeSector.named.StructName())
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: initialize sector failed, pkey not found", runtimeSector.named.StructName())
	} // if

	runtimeSector.layoutJson = layoutJson
	runtimeSector.layoutType = layoutType
	return nil
}
