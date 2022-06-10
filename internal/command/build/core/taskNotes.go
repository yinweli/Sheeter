package core

import "fmt"

// executeExcel 建立欄位註解
func (this *Task) executeNotes() error {
	notes := this.getCols(this.global.LineOfNote)

	if notes == nil {
		return fmt.Errorf("note line not found: %s", this.logName())
	} // if

	count := len(notes)

	for col, itor := range this.columns {
		if col >= 0 && col < count { // 註解行的數量可能因為空白格的關係會短缺, 所以要檢查一下
			itor.Note = notes[col]
		} // if
	} // for

	return nil
}
