package core

import (
	"fmt"

	"github.com/emirpasic/gods/sets/hashset"
)

// TaskFields 建立欄位列表
func TaskFields(ctx *Context) error {
	fields := ctx.GetCols(ctx.Global.LineOfField)

	if fields == nil {
		return fmt.Errorf("fields not found: %s", ctx.LogName())
	} // if

	names := hashset.New()
	pkey := false
	ctx.Columns = []*Column{} // 把欄位列表清空, 避免不必要的問題

	for _, itor := range fields {
		if len(itor) <= 0 {
			break
		} // if

		name, field, err := ParseField(itor)

		if err != nil {
			return fmt.Errorf("field parser failed: %s [%s : %s]", ctx.LogName(), itor, err)
		} // if

		if names.Contains(name) {
			return fmt.Errorf("name duplicate: %s [%s]", ctx.LogName(), itor)
		} // if

		names.Add(name)

		if field.IsPkey() && pkey {
			return fmt.Errorf("pkey duplicate: %s [%s]", ctx.LogName(), itor)
		} // if

		if field.IsPkey() {
			pkey = true
		} // if

		ctx.Columns = append(ctx.Columns, &Column{
			Name:  name,
			Field: field,
		})
	} // for

	if pkey == false { // 這裡其實也順便檢查了沒有欄位的問題
		return fmt.Errorf("pkey not found: %s", ctx.LogName())
	} // if

	return nil
}
