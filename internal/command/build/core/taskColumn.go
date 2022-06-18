package core

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

// executeColumn 建立欄位列表
func (this *Task) executeColumn() error {
	fields := this.getRowContent(this.global.LineOfField)

	if fields == nil {
		return fmt.Errorf("field line not found: %s", this.originalName())
	} // if

	notes := this.getRowContent(this.global.LineOfNote)

	if notes == nil {
		return fmt.Errorf("note line not found: %s", this.originalName())
	} // if

	names := hashset.New()
	pkey := false
	this.columns = []*Column{} // 把欄位列表清空, 避免不必要的問題

	for col, itor := range fields {
		if len(itor) <= 0 { // 一旦遇到空格, 就結束建立欄位列表
			break
		} // if

		name, field, err := ParseField(itor)

		if err != nil {
			return fmt.Errorf("field parser failed: %s [%s : %s]", this.originalName(), itor, err)
		} // if

		if names.Contains(name) {
			return fmt.Errorf("name duplicate: %s [%s]", this.originalName(), itor)
		} // if

		names.Add(name)

		if field.IsPkey() && pkey {
			return fmt.Errorf("pkey duplicate: %s [%s]", this.originalName(), itor)
		} // if

		if field.IsPkey() {
			pkey = true
		} // if

		note := ""

		if col >= 0 && col < len(notes) { // 註解行的數量可能因為空白格的關係會短缺, 所以要檢查一下
			note = notes[col]
		} // if

		this.columns = append(this.columns, &Column{
			Name:  name,
			Note:  note,
			Field: field,
		})
	} // for

	if pkey == false { // 這裡其實也順便檢查了沒有欄位的問題
		return fmt.Errorf("pkey not found: %s", this.originalName())
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
