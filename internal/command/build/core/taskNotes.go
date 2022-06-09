package core

import (
	"fmt"
)

// TaskNotes 建立欄位註解
func TaskNotes(ctx *Context) error {
	notes := ctx.GetCols(ctx.Global.LineOfNote)

	if notes == nil {
		return fmt.Errorf("note line not found: %s", ctx.LogName())
	} // if

	count := len(notes)

	for col, itor := range ctx.Columns {
		if col >= 0 && col < count { // 註解行的數量可能因為空白格的關係會短缺, 所以要檢查一下
			itor.Note = notes[col]
		} // if
	} // for

	return nil
}
