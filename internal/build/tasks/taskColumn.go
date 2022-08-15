package tasks

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/internal/build/layers"
	"github.com/yinweli/Sheeter/internal/build/layouts"
)

// column 建立欄位列表
func (this *Task) column() error {
	fieldLine, err := this.getRowContent(this.LineOfField)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nfield line not found", this.originalName())
	} // if

	layerLine, err := this.getRowContent(this.LineOfLayer)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nlayer line not found", this.originalName())
	} // if

	noteLine, err := this.getRowContent(this.LineOfNote)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nnote line not found", this.originalName())
	} // if

	this.builder = layouts.NewBuilder()

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.Parser(itor)

		if err != nil {
			return fmt.Errorf("read column failed: %s [%s]\nfield parser failed\n%w", this.originalName(), itor, err)
		} // if

		layer, back, err := layers.Parser(catch(layerLine, col))

		if err != nil {
			return fmt.Errorf("read column failed: %s [%s]\nlayer parser failed\n%w", this.originalName(), itor,
				err)
		} // if

		note := catch(noteLine, col)

		if err := this.builder.Add(name, note, field, layer, back); err != nil {
			return fmt.Errorf("read column failed: %s [%s]\nadd to builder failed\n%w", this.originalName(), itor,
				err)
		} // if
	} // for

	pkeyCount := this.builder.PkeyCount()

	if pkeyCount > 1 {
		return fmt.Errorf("read column failed: %s\npkey duplicate", this.originalName())
	} // if

	if pkeyCount <= 0 {
		return fmt.Errorf("read column failed: %s\npkey not found", this.originalName())
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}

// catch 從列表中取得項目
func catch(lists []string, index int) string {
	if index >= 0 && index < len(lists) { // 列表的數量可能因為空白格的關係會短缺, 所以要檢查一下
		return lists[index]
	} // if

	return ""
}
