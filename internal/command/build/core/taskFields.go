package core

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

// executeFields 建立欄位列表
func (this *Task) executeFields() error {
	fields := this.getCols(this.global.LineOfField)

	if fields == nil {
		return fmt.Errorf("fields not found: %s", this.logName())
	} // if

	names := hashset.New()
	pkey := false
	this.columns = []*Column{} // 把欄位列表清空, 避免不必要的問題

	for _, itor := range fields {
		if len(itor) <= 0 { // 一旦遇到空格, 就結束建立欄位列表
			break
		} // if

		name, field, err := ParseField(itor)

		if err != nil {
			return fmt.Errorf("field parser failed: %s [%s : %s]", this.logName(), itor, err)
		} // if

		if names.Contains(name) {
			return fmt.Errorf("name duplicate: %s [%s]", this.logName(), itor)
		} // if

		names.Add(name)

		if field.IsPkey() && pkey {
			return fmt.Errorf("pkey duplicate: %s [%s]", this.logName(), itor)
		} // if

		if field.IsPkey() {
			pkey = true
		} // if

		this.columns = append(this.columns, &Column{
			Name:  name,
			Field: field,
		})
	} // for

	if pkey == false { // 這裡其實也順便檢查了沒有欄位的問題
		return fmt.Errorf("pkey not found: %s", this.logName())
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressS)
	} // if

	return nil
}
