package core

import "fmt"

// TaskColumns 建立欄位列表
func TaskColumns(ctx *Context) error {
	lineOfField := ctx.Global.GetLineOfField()

	if len(ctx.Sheets) <= lineOfField {
		return fmt.Errorf("field line not found: %s", ctx.LogName())
	} // if

	fields := ctx.Sheets[lineOfField]
	parser := NewParser()
	ctx.Columns = []*Column{} // 把行資料列表清空, 避免不必要的問題

	for _, itor := range fields {
		if len(itor) <= 0 {
			break
		} // if

		name, field, err := parser.Parse(itor)

		if err != nil {
			return fmt.Errorf("parse failed: %s [%s : %s]", ctx.LogName(), itor, err)
		} // if

		if field.IsPkey() && ctx.Pkey != nil {
			return fmt.Errorf("too many pkey: %s", ctx.LogName())
		} // if

		column := &Column{
			Name:  name,
			Field: field,
		}

		if field.IsPkey() {
			ctx.Pkey = column
		} // if

		ctx.Columns = append(ctx.Columns, column)
	} // for

	if ctx.Pkey == nil { // 這裡檢查是否有主要索引, 也同時檢查了沒有任何欄位的情況
		return fmt.Errorf("pkey not found: %s", ctx.LogName())
	} // if

	_ = ctx.Progress.Add(1)
	return nil
}
