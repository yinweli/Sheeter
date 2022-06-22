package core

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

// runColumn 建立欄位列表
func (this *Task) runColumn() error {
	fields, err := this.getRowContent(this.global.LineOfField)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nfield line not found", this.originalName())
	} // if

	notes, err := this.getRowContent(this.global.LineOfNote)

	if err != nil {
		return fmt.Errorf("read column failed: %s\nnote line not found", this.originalName())
	} // if

	this.columns = []*Column{} // 把欄位列表清空, 避免不必要的問題
	names := hashset.New()
	pkey := false

	for col, itor := range fields {
		if len(itor) <= 0 { // 一旦遇到空格, 就結束建立欄位列表
			break
		} // if

		name, field, err := ParseField(itor)

		if err != nil {
			return fmt.Errorf("read column failed: %s [%s]\nfield parser faile\n%s", this.originalName(), itor, err)
		} // if

		if names.Contains(name) {
			return fmt.Errorf("read column failed: %s [%s]\nname duplicate", this.originalName(), itor)
		} // if

		names.Add(name)

		if field.IsPkey() && pkey {
			return fmt.Errorf("read column failed: %s [%s]\npkey duplicate", this.originalName(), itor)
		} // if

		if field.IsPkey() {
			pkey = true
		} // if

		note := ""

		if col >= 0 && col < len(notes) { // 註解的數量可能因為空白格的關係會短缺, 所以要檢查一下
			note = notes[col]
		} // if

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
