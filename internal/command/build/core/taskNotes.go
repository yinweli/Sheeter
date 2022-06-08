package core

import "fmt"

// TaskNotes 建立欄位註解
func TaskNotes(ctx *Context) error {
	lineOfNote := ctx.Global.GetLineOfNote()

	if len(ctx.Sheets) <= lineOfNote {
		return fmt.Errorf("note line not found: %s", ctx.LogName())
	} // if

	notes := ctx.Sheets[lineOfNote]

	for col, itor := range ctx.Columns {
		var data string

		if col >= 0 && col < len(notes) { // 註解行的數量可能因為空白格的關係會短缺, 所以要檢查一下
			data = notes[col]
		} // if

		itor.Note = data
	} // for

	_ = ctx.Progress.Add(1)
	return nil
}
