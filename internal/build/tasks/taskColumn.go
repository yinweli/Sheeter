package tasks

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/internal/build/layouts"
)

// runColumn 建立欄位列表
func (this *Task) runColumn() error {
	fieldLine, err := this.getRowContent(this.global.LineOfField)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nfield line not found", this.originalName())
	} // if

	layerLine, err := this.getRowContent(this.global.LineOfLayer)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nlayer line not found", this.originalName())
	} // if

	noteLine, err := this.getRowContent(this.global.LineOfNote)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nnote line not found", this.originalName())
	} // if

	this.columns = []*Column{} // 把欄位列表清空, 避免不必要的問題
	duplField := layouts.NewDuplField()
	duplLayer := layouts.NewDuplLayer()
	pkey := false

	for col, itor := range fieldLine {
		if itor == "" { // 一旦遇到空欄位, 就結束建立欄位列表
			break
		} // if

		name, field, err := fields.ParseField(itor)

		if err != nil {
			return fmt.Errorf("read column failed: %s [%s]\nfield parser failed\n%w", this.originalName(), itor, err)
		} // if

		if duplField.Check(name) == false {
			return fmt.Errorf("read column failed: %s [%s]\nfield duplicate", this.originalName(), itor)
		} // if

		if field.IsPkey() && pkey { // 只能有一個主要索引
			return fmt.Errorf("read column failed: %s [%s]\npkey duplicate", this.originalName(), itor)
		} // if

		if field.IsPkey() {
			pkey = true
		} // if

		layer := fromList(layerLine, col)
		layers, _, err := layers2.ParseLayer(layer)

		if err != nil {
			return fmt.Errorf("read column failed: %s [%s]\nlayer parser failed\n%w", this.originalName(), itor,
				err)
		} // if

		if duplLayer.Check(layers...) == false {
			return fmt.Errorf("read column failed: %s [%s]\nfield duplicate", this.originalName(), itor)
		} // if

		note := fromList(noteLine, col)
		this.columns = append(this.columns, &Column{
			Name:  name,
			Note:  note,
			Field: field,
		})
	} // for

	if pkey == false { // 這裡其實也順便檢查了沒有欄位的問題
		return fmt.Errorf("read column failed: %s\npkey not found", this.originalName())
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}

// fromList 從列表中取得項目
func fromList(lists []string, index int) string {
	if index >= 0 && index < len(lists) { // 列表的數量可能因為空白格的關係會短缺, 所以要檢查一下
		return lists[index]
	} // if

	return ""
}
