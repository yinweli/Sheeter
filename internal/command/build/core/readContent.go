package core

import (
	"fmt"
)

// ReadContent 讀取內容
func ReadContent(cargo *Cargo) error {
	pkey, err := buildColumns(cargo)

	if err != nil {
		return err
	} // if

	err = buildNotes(cargo)

	if err != nil {
		return err
	} // if

	err = buildDatas(cargo) // 目前buildDatas不會失敗
	err = pkeyCheck(cargo, pkey)

	if err != nil {
		return err
	} // if

	return nil
}

// buildColumns 建立行資料列表
func buildColumns(cargo *Cargo) (pkey *Column, err error) {
	lineOfField := cargo.Global.GetLineOfField()

	if len(cargo.Sheets) <= lineOfField {
		return nil, fmt.Errorf("field line not found: %s", cargo.LogName())
	} // if

	fields := cargo.Sheets[lineOfField]
	parser := NewParser()
	cargo.Columns = []*Column{} // 把行資料列表清空, 避免不必要的問題

	for _, itor := range fields {
		if len(itor) <= 0 {
			break
		} // if

		name, field, err := parser.Parse(itor)

		if err != nil {
			return nil, fmt.Errorf("field parse failed: %s [%s : %s]", cargo.LogName(), itor, err)
		} // if

		if field.PrimaryKey() && pkey != nil {
			return nil, fmt.Errorf("too many pkey: %s", cargo.LogName())
		} // if

		column := &Column{
			Name:  name,
			Field: field,
		}

		if field.PrimaryKey() {
			pkey = column
		} // if

		cargo.Columns = append(cargo.Columns, column)
	} // for

	if pkey == nil { // 這裡檢查是否有主要索引, 也同時檢查了沒有任何欄位的情況
		return nil, fmt.Errorf("must have one pkey: %s", cargo.LogName())
	} // if

	return pkey, nil
}

// buildNotes 建立欄位註解
func buildNotes(cargo *Cargo) error {
	lineOfNote := cargo.Global.GetLineOfNote()

	if len(cargo.Sheets) <= lineOfNote {
		return fmt.Errorf("note line not found: %s", cargo.LogName())
	} // if

	notes := cargo.Sheets[lineOfNote]

	for col, itor := range cargo.Columns {
		var data string

		if col >= 0 && col < len(notes) { // 註解行的數量可能因為空白格的關係會短缺, 所以要檢查一下
			data = notes[col]
		} // if

		itor.Note = data
	} // for

	return nil
}

// buildDatas 建立資料
func buildDatas(cargo *Cargo) error {
	for _, itor := range cargo.Columns {
		itor.Datas = []string{} // 把資料列表清空, 避免不必要的問題
	} // for

	for row := cargo.Global.GetLineOfData(); row < len(cargo.Sheets); row++ {
		datas := cargo.Sheets[row]

		for col, itor := range cargo.Columns {
			var data string

			if col >= 0 && col < len(datas) { // 資料行的數量可能因為空白格的關係會短缺, 所以要檢查一下
				data = datas[col]
			} // if

			cargo.Progress.Add(1)
			itor.Datas = append(itor.Datas, data)
		} // for
	} // for

	return nil
}

// pkeyCheck 主要索引檢查
func pkeyCheck(cargo *Cargo, pkey *Column) error {
	datas := make(map[string]bool)

	for _, itor := range pkey.Datas {
		if _, exist := datas[itor]; exist {
			return fmt.Errorf("pkey duplicate: %s", cargo.LogName())
		} // if

		datas[itor] = true
	} // for

	return nil
}
