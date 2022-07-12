package core

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// runColumn 建立欄位列表
func (this *Task) runColumn() error {
	fieldLine, err := this.getRowContent(this.global.LineOfField)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nfield line not found", this.originalName())
	} // if

	noteLine, err := this.getRowContent(this.global.LineOfNote)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nnote line not found", this.originalName())
	} // if

	this.columns = []*Column{} // 把欄位列表清空, 避免不必要的問題
	duplicateField := util.Duplicate{}
	pkey := false

	for col, itor := range fieldLine {
		if len(itor) <= 0 { // 一旦遇到空格, 就結束建立欄位列表
			break
		} // if

		name, field, err := ParseField(itor)

		if err != nil {
			return fmt.Errorf("read column failed: %s [%s]\nfield parser failed\n%w", this.originalName(), itor, err)
		} // if

		if duplicateField.Check(name) == false { // 欄位名稱不可重複
			return fmt.Errorf("read column failed: %s [%s]\nfield duplicate", this.originalName(), itor)
		} // if

		if field.IsPkey() && pkey { // 只能有一個主要索引
			return fmt.Errorf("read column failed: %s [%s]\npkey duplicate", this.originalName(), itor)
		} // if

		if field.IsPkey() {
			pkey = true
		} // if

		note := column(noteLine, col)

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

// column 從欄位列表中取得欄位
func column(columns []string, col int) string {
	if col >= 0 && col < len(columns) { // 欄位的數量可能因為空白格的關係會短缺, 所以要檢查一下
		return columns[col]
	} // if

	return ""
}
