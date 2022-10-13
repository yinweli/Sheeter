package excels

import (
	"fmt"
	"sort"

	"github.com/xuri/excelize/v2"
)

// Excel excel資料
type Excel struct {
	excel *excelize.File // excel物件
}

// Open 開啟excel
func (this *Excel) Open(name string) error {
	excel, err := excelize.OpenFile(name)

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

// Get 取得表格
func (this *Excel) Get(name string) (sheet *Sheet, err error) {
	rows, err := this.excel.Rows(name)

	if err != nil {
		return nil, fmt.Errorf("get failed: %w", err)
	} // if

	return &Sheet{rows: rows}, nil
}

// GetLine 取得表格行資料
func (this *Excel) GetLine(name string, line ...int) (result map[int][]string, err error) {
	sheet, err := this.Get(name)

	if err != nil {
		return nil, fmt.Errorf("get line failed: %w", err)
	} // if

	defer sheet.Close()
	result = map[int][]string{}
	current := 0 // 最少要一次才能定位到第1行, 所以起始位置設為0
	sort.Ints(line)

	for _, itor := range line {
		if itor <= 0 { // 最少要一次才能定位到第1行, 所以若起始位置設為0line <= 0, 就表示錯誤
			return nil, fmt.Errorf("get line failed: line <= 0")
		} // if

		if sheet.Nextn(itor-current) == false {
			return nil, fmt.Errorf("get line failed: line not found")
		} // if

		current = itor
		data, err := sheet.Data()

		if err != nil {
			return nil, fmt.Errorf("get line failed: %w", err)
		} // if

		if data == nil { // 如果取得空行, 就回傳個空切片吧
			data = []string{}
		} // if

		result[itor] = data
	} // for

	return result, nil
}

// Exist 表格是否存在
func (this *Excel) Exist(sheet string) bool {
	return this.excel.GetSheetIndex(sheet) != -1
}

// IsOpen 是否開啟excel
func (this *Excel) IsOpen() bool {
	return this.excel != nil
}

// Sheet 表格資料
type Sheet struct {
	rows *excelize.Rows
}

// Close 關閉表格資料
func (this *Sheet) Close() {
	if this.rows != nil {
		_ = this.rows.Close()
		this.rows = nil
	} // if
}

// Next 往下一行
func (this *Sheet) Next() bool {
	return this.rows.Next()
}

// Nextn 往下n行
func (this *Sheet) Nextn(n int) bool {
	if n < 0 {
		return false
	} // if

	for i := 0; i < n; i++ {
		if this.rows.Next() == false {
			return false
		} // if
	} // for

	return true
}

// Data 取得行資料
func (this *Sheet) Data() (result []string, err error) {
	result, err = this.rows.Columns()

	if err != nil {
		return nil, fmt.Errorf("data failed: %w", err)
	} // if

	return result, nil
}
