package excels

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Excel excel資料
type Excel struct {
	excel *excelize.File // excel物件
}

// Open 開啟excel
func (this *Excel) Open(filename string) error {
	excel, err := excelize.OpenFile(filename)

	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	} // if

	this.excel = excel
	return nil
}

// Close 關閉excel
func (this *Excel) Close() {
	if this.excel != nil {
		_ = this.excel.Close()
		this.excel = nil
	} // if
}

// GetLine 取得excel行資料, index從1起算; 當行不存在時不會發生錯誤
func (this *Excel) GetLine(sheet string, index int) (line *Line, err error) {
	if index <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("get line failed: index <= 0")
	} // if

	rows, err := this.excel.Rows(sheet)

	if err != nil {
		return nil, fmt.Errorf("get line failed: %w", err)
	} // if

	for l := 0; l < index; l++ {
		rows.Next()
	} // for

	return &Line{rows: rows}, nil
}

// GetData 取得excel行資料列表, index從1起算; 當行不存在時會發生錯誤
func (this *Excel) GetData(sheet string, index int) (data []string, err error) {
	if index <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("get data failed: index <= 0")
	} // if

	rows, err := this.excel.Rows(sheet)

	if err != nil {
		return nil, fmt.Errorf("get data failed: %w", err)
	} // if

	defer func() { _ = rows.Close() }()

	for l := 0; l < index; l++ {
		if rows.Next() == false {
			return nil, fmt.Errorf("get data failed: data not found")
		} // if
	} // for

	data, err = rows.Columns()

	if err != nil {
		return nil, fmt.Errorf("get data failed: invalid columns: %w", err)
	} // if

	if data == nil {
		data = []string{} // 如果取得空行, 就回傳個空切片吧
	} // if

	return data, nil
}

// SheetExist 表單是否存在
func (this *Excel) SheetExist(sheet string) bool {
	return this.excel.GetSheetIndex(sheet) != -1
}

// IsOpen 是否開啟excel
func (this *Excel) IsOpen() bool {
	return this.excel != nil
}

// Line excel行資料
type Line struct {
	rows *excelize.Rows
}

// Close 關閉行資料
func (this *Line) Close() {
	if this.rows != nil {
		_ = this.rows.Close()
		this.rows = nil
	} // if
}

// Next 往下一行
func (this *Line) Next() bool {
	return this.rows.Next()
}

// Data 取得資料列表
func (this *Line) Data() (result []string, err error) {
	result, err = this.rows.Columns()

	if err != nil {
		return nil, fmt.Errorf("data failed: %w", err)
	} // if

	return result, nil
}
