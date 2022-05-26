package core

import (
	"fmt"
	"path"

	"Sheeter/internal/util"

	"github.com/xuri/excelize/v2"
)

// ReadSheet 讀取表格
func ReadSheet(cargo *Cargo, task int) error {
	sheet, err := buildSheet(cargo, task)

	if err != nil {
		return err
	} // if

	pkey, err := buildColumns(cargo, sheet)

	if err != nil {
		return err
	} // if

	_ = buildNotes(cargo, sheet) // 目前buildNotes不會失敗
	_ = buildDatas(cargo, sheet) // 目前buildDatas不會失敗
	err = pkeyCheck(cargo, pkey)

	if err != nil {
		return err
	} // if

	return nil
}

// buildSheet 建立表格列表
func buildSheet(cargo *Cargo, task int) (result sheets, err error) {
	file, err := excelize.OpenFile(path.Join(cargo.Global.ExcelPath, cargo.Element.Excel))

	if err != nil {
		return nil, err
	} // if

	defer func() {
		_ = file.Close()
	}()

	result, err = file.GetRows(cargo.Element.Sheet)

	if result == nil || err != nil {
		return nil, fmt.Errorf("sheet not found: %s", cargo.Element.GetFullName())
	} // if

	if len(result) < 2 { // 表格最少要有2行: 註解行, 欄位行
		return nil, fmt.Errorf("sheet have too less line: %s", cargo.Element.GetFullName())
	} // if

	cargo.Progress.ChangeMax(util.CalcProgress(len(result), task))

	return result, nil
}

// buildColumns 建立行資料列表
func buildColumns(cargo *Cargo, sheet sheets) (pkey *Column, err error) {
	fields := sheet[cargo.Global.GetLineOfField()]
	parser := NewParser()
	cargo.Columns = []*Column{} // 把行資料列表清空, 避免不必要的問題

	for _, itor := range fields {
		if len(itor) <= 0 {
			break
		} // if

		name, field, err := parser.Parse(itor)

		if err != nil {
			return nil, fmt.Errorf("field parse failed: %s [%s : %s]", cargo.Element.GetFullName(), itor, err)
		} // if

		if field.PrimaryKey() && pkey != nil {
			return nil, fmt.Errorf("too many pkey: %s", cargo.Element.GetFullName())
		} // if

		column := &Column{
			Name:  name,
			Field: field,
		}

		if field.PrimaryKey() {
			pkey = column
		} // if

		cargo.Columns = append(cargo.Columns, column)
		_ = cargo.Progress.Add(1)
	} // for

	if pkey == nil { // 這裡也同時檢查了沒有任何欄位的情況
		return nil, fmt.Errorf("must have one pkey: %s", cargo.Element.GetFullName())
	} // if

	return pkey, nil
}

// buildNotes 建立欄位註解
func buildNotes(cargo *Cargo, sheet sheets) error {
	notes := sheet[cargo.Global.GetLineOfNote()]

	for col, itor := range cargo.Columns {
		var data string

		if col >= 0 && col < len(notes) { // 註解行的數量可能因為空白格的關係會短缺, 所以要檢查一下
			data = notes[col]
		} // if

		itor.Note = data
		_ = cargo.Progress.Add(1)
	} // for

	return nil
}

// buildDatas 建立資料
func buildDatas(cargo *Cargo, sheet sheets) error {
	for _, itor := range cargo.Columns {
		itor.Datas = []string{} // 把資料列表清空, 避免不必要的問題
	} // for

	for row := cargo.Global.GetLineOfData(); row < len(sheet); row++ {
		datas := sheet[row]

		for col, itor := range cargo.Columns {
			var data string

			if col >= 0 && col < len(datas) { // 資料行的數量可能因為空白格的關係會短缺, 所以要檢查一下
				data = datas[col]
			} // if

			itor.Datas = append(itor.Datas, data)
			_ = cargo.Progress.Add(1)
		} // for
	} // for

	return nil
}

// pkeyCheck 主要索引檢查
func pkeyCheck(cargo *Cargo, pkey *Column) error {
	datas := make(map[string]bool)

	for _, itor := range pkey.Datas {
		if _, exist := datas[itor]; exist {
			return fmt.Errorf("pkey duplicate: %s", cargo.Element.GetFullName())
		} // if

		datas[itor] = true
	} // for

	return nil
}

// sheets 表格資料列表
type sheets [][]string
