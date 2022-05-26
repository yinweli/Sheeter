package core

import (
	"fmt"
	"path"

	"Sheeter/internal"

	"github.com/xuri/excelize/v2"
)

// ReadSheet 讀取表格
func ReadSheet(task *Task) error {
	file, err := excelize.OpenFile(path.Join(task.Global.ExcelPath, task.Element.Excel))

	if err != nil {
		return err
	} // if

	defer func() {
		_ = file.Close()
	}()

	sheet, err := file.GetRows(task.Element.Sheet) // 從表格中取得全部資料 TODO: 分離子函式

	if sheet == nil || err != nil {
		return fmt.Errorf("sheet not found: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	if len(sheet) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return fmt.Errorf("sheet have too less line: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	task.Progress.ChangeMax(len(sheet) * internal.TaskMax) // 設定進度條最大值

	// 讀取欄位名稱, 欄位類型 TODO: 分離子函式
	fields := sheet[task.Global.GetLineOfField()]
	parser := NewParser()
	pkeyCount := 0

	for _, itor := range fields {
		if len(itor) <= 0 {
			break
		} // if

		name, field, err := parser.Parse(itor)

		if err != nil {
			return fmt.Errorf("sheet field Parse failed: %s(%s) [%s]", task.Element.Excel, task.Element.Sheet, err)
		} // if

		if field.PrimaryKey() {
			pkeyCount++
		} // if

		column := &Column{
			Name:  name,
			Field: field,
		}
		task.Columns = append(task.Columns, column)
		_ = task.Progress.Add(1)
	} // for

	if len(task.Columns) <= 0 { // 表格最少要有一個欄位
		return fmt.Errorf("sheet have too less field: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	if pkeyCount != 0 { // 表格必須有一個主要索引, 超過也不行
		return fmt.Errorf("sheet must, and only have one pkey: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	// 讀取欄位註解 TODO: 分離子函式
	notes := sheet[task.Global.GetLineOfNote()]

	for col := 0; col < len(task.Columns); col++ {
		var data string

		if col >= 0 && col < len(notes) {
			data = notes[col]
		} // if

		task.Columns[col].Note = data
		_ = task.Progress.Add(1)
	} // for

	// 讀取資料 TODO: 分離子函式
	for row := task.Global.GetLineOfData(); row < len(sheet); row++ {
		datas := sheet[row]

		for col := 0; col < len(task.Columns); col++ {
			var data string

			if col >= 0 && col < len(datas) {
				data = datas[col]
			} // if

			task.Columns[col].Datas = append(task.Columns[col].Datas, data)
			_ = task.Progress.Add(1)
		} // for
	} // for

	// 檢查重複索引 TODO: 分離子函式
	/*
		pkeys := mapset.NewSet[string]()
		duplicate := ""

		for _, itor := range task.Columns {
			if itor.Field.PrimaryKey() {
				for _, pkey := range itor.Datas {
					if pkeys.Add(pkey) == false {
						duplicate = strings.Join([]string{duplicate, pkey}, internal.ArraySeparator)
					} // if
				} // for

				break
			} // if
		} // for

		if duplicate != "" {
			return fmt.Errorf("pkey duplicate: %s(%s) [%s]", task.Element.Excel, task.Element.Sheet, duplicate)
		} // if
	*/

	return nil
}

// 獲得表格資料
func getSheet(task *Task) (sheet [][]string, err error) {
	file, err := excelize.OpenFile(path.Join(task.Global.ExcelPath, task.Element.Excel))

	if err != nil {
		return nil, err
	} // if

	defer func() {
		_ = file.Close()
	}()

	sheet, err = file.GetRows(task.Element.Sheet) // 從表格中取得全部資料 TODO: 分離子函式

	if sheet == nil || err != nil {
		return nil, fmt.Errorf("sheet not found: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	if len(sheet) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return nil, fmt.Errorf("sheet have too less line: %s(%s)", task.Element.Excel, task.Element.Sheet)
	} // if

	return sheet, nil
}
