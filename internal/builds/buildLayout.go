package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/builds/fields"
	"github.com/yinweli/Sheeter/internal/builds/layers"
	"github.com/yinweli/Sheeter/internal/builds/layouts"
	"github.com/yinweli/Sheeter/internal/util"
)

// buildLayout 建立布局資料
func buildLayout(content *Content) error {
	fieldLine, err := content.getColumns(content.LineOfField)

	if err != nil {
		return fmt.Errorf("%s: build layout failed, field line not found", content.TargetName())
	} // if

	layerLine, err := content.getColumns(content.LineOfLayer)

	if err != nil {
		return fmt.Errorf("%s: build layout failed, layer line not found", content.TargetName())
	} // if

	noteLine, err := content.getColumns(content.LineOfNote)

	if err != nil {
		return fmt.Errorf("%s: build layout failed, note line not found", content.TargetName())
	} // if

	builder := layouts.NewBuilder()

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("%s: build layout failed: %w", content.TargetName(), err)
		} // if

		layer, back, err := layers.Parser(util.GetItem(layerLine, col))

		if err != nil {
			return fmt.Errorf("%s: build layout failed: %w", content.TargetName(), err)
		} // if

		note := util.GetItem(noteLine, col)

		if err := builder.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("%s: build layout failed: %w", content.TargetName(), err)
		} // if
	} // for

	pkeyCount := builder.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("%s: build layout failed, pkey duplicate", content.TargetName())
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("%s: build layout failed, pkey not found", content.TargetName())
	} // if

	content.builder = builder
	return nil
}
