package core

import (
	"fmt"
	"path"

	"Sheeter/internal"

	"github.com/xuri/excelize/v2"
)

// ReadSheet 讀取表格
func ReadSheet(task *Task) error {
	sheet, err := buildSheet(task)

	if err != nil {
		return err
	} // if

	pkey, err := buildColumns(task, sheet)

	if err != nil {
		return err
	} // if

	_ = buildNotes(task, sheet) // 目前buildNotes不會失敗
	_ = buildDatas(task, sheet) // 目前buildDatas不會失敗
	err = pkeyCheck(task, pkey)

	if err != nil {
		return err
	} // if

	return nil
}

// buildSheet 建立表格列表
func buildSheet(task *Task) (sheet [][]string, err error) {
	file, err := excelize.OpenFile(path.Join(task.Global.ExcelPath, task.Element.Excel))

	if err != nil {
		return nil, err
	} // if

	defer func() {
		_ = file.Close()
	}()

	sheet, err = file.GetRows(task.Element.Sheet)

	if sheet == nil || err != nil {
		return nil, fmt.Errorf("sheet not found: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	if len(sheet) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return nil, fmt.Errorf("sheet have too less line: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	task.Progress.ChangeMax(len(sheet) * internal.TaskMax) // 設定進度條最大值

	return sheet, nil
}

// buildColumns 建立行資料列表
func buildColumns(task *Task, sheet [][]string) (pkey *Column, err error) {
	fields := sheet[task.Global.GetLineOfField()]
	parser := NewParser()
	task.Columns = []*Column{} // 把行資料列表清空, 避免不必要的問題

	for _, itor := range fields {
		if len(itor) <= 0 {
			break
		} // if

		name, field, err := parser.Parse(itor)

		if err != nil {
			return nil, fmt.Errorf("field parse failed: %s(%s) [%s : %s]", task.Element.Excel, task.Element.Sheet, itor, err)
		} // if

		if field.PrimaryKey() && pkey != nil {
			return nil, fmt.Errorf("too many pkey: %s(%s)", task.Element.Excel, task.Element.Sheet)
		} // if

		column := &Column{
			Name:  name,
			Field: field,
		}

		if field.PrimaryKey() {
			pkey = column
		} // if

		task.Columns = append(task.Columns, column)
		_ = task.Progress.Add(1)
	} // for

	if pkey == nil { // 這裡也同時檢查了沒有任何欄位的情況
		return nil, fmt.Errorf("must have one pkey: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	return pkey, nil
}

// buildNotes 建立欄位註解
func buildNotes(task *Task, sheet [][]string) error {
	notes := sheet[task.Global.GetLineOfNote()]

	for pos, itor := range task.Columns {
		var data string

		if pos >= 0 && pos < len(notes) { // 註解行的數量可能因為空白格的關係會短缺, 所以要檢查一下
			data = notes[pos]
		} // if

		itor.Note = data
		_ = task.Progress.Add(1)
	} // for

	return nil
}

// buildDatas 建立資料
func buildDatas(task *Task, sheet [][]string) error {
	for _, itor := range task.Columns {
		itor.Datas = []string{} // 把資料列表清空, 避免不必要的問題
	} // for

	for row := task.Global.GetLineOfData(); row < len(sheet); row++ {
		datas := sheet[row]

		for pos, itor := range task.Columns {
			var data string

			if pos >= 0 && pos < len(datas) { // 資料行的數量可能因為空白格的關係會短缺, 所以要檢查一下
				data = datas[pos]
			} // if

			itor.Datas = append(itor.Datas, data)
			_ = task.Progress.Add(1)
		} // for
	} // for

	return nil
}

// pkeyCheck 主要索引檢查
func pkeyCheck(task *Task, pkey *Column) error {
	datas := make(map[string]bool)

	for _, itor := range pkey.Datas {
		if _, exist := datas[itor]; exist {
			return fmt.Errorf("pkey duplicate: %s(%s)", task.Element.Excel, task.Element.Sheet)
		} // if

		datas[itor] = true
	} // for

	return nil
}
