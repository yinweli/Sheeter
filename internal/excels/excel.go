package excels

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/thedatashed/xlsxreader"
)

// Excel excel資料
type Excel struct {
	excel *xlsxreader.XlsxFileCloser // excel物件
}

// Open 開啟excel
func (this *Excel) Open(name string) error {
	excel, err := xlsxreader.OpenFile(name)

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

// Get 取得表單
func (this *Excel) Get(name string) (sheet *Sheet, err error) {
	if this.Exist(name) == false {
		return nil, fmt.Errorf("get failed: sheet not exist")
	} // if

	sheet = &Sheet{
		rows: this.excel.ReadRows(name),
	}
	return sheet, nil
}

// GetLine 取得表單行資料
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

		data := []string{}

		if sheet.Nextn(itor - current) {
			current = itor

			if data, err = sheet.Data(); err != nil {
				return nil, fmt.Errorf("get line failed: %w", err)
			} // if

			if data == nil { // 如果取得空行, 就回傳個空切片吧
				data = []string{}
			} // if
		} // if

		result[itor] = data
	} // for

	return result, nil
}

// Sheets 取得表單列表
func (this *Excel) Sheets() []string {
	if this.excel != nil {
		return this.excel.Sheets
	} // if

	return []string{}
}

// Exist 表單是否存在
func (this *Excel) Exist(name string) bool {
	if this.excel != nil {
		for _, itor := range this.excel.Sheets {
			if itor == name {
				return true
			} // if
		} // for
	} // if

	return false
}

// IsOpen 是否開啟excel
func (this *Excel) IsOpen() bool {
	return this.excel != nil
}

// Sheet 表單資料
type Sheet struct {
	rows chan xlsxreader.Row // 表單資料
	row  *xlsxreader.Row     // 行資料
	line int                 // 目前行數
}

// Close 關閉表單資料
func (this *Sheet) Close() {
	if this.rows != nil {
		// 由於xlsxreader的要求, 必須在關閉前把表單尋訪完畢
		// 不然會遺留未完成的goroutine與channel物件
		for range this.rows {
			// do nothing...
		} // for

		this.rows = nil
	} // if
}

// Next 往下一行
func (this *Sheet) Next() bool {
	if this.row != nil && this.row.Index != this.line {
		this.line++
		return true
	} // if

	row := <-this.rows
	this.row = &row
	this.line++
	return this.row.Error == nil && this.row.Index > 0 && this.row.Cells != nil
}

// Nextn 往下n行
func (this *Sheet) Nextn(n int) bool {
	if n < 0 {
		return false
	} // if

	for i := 0; i < n; i++ {
		if this.Next() == false {
			return false
		} // if
	} // for

	return true
}

// Data 取得行資料
func (this *Sheet) Data() (result []string, err error) {
	if this.row == nil {
		return nil, fmt.Errorf("data failed: row nil")
	} // if

	if this.row.Index != this.line {
		return nil, nil
	} // if

	for _, itor := range this.row.Cells {
		index := columnToIndex(itor.Column)

		for len(result) < index {
			result = append(result, "")
		} // for

		result[index-1] = itor.Value // 由於欄位從1開始, 而陣列從0開始, 所以要減1
	} // for

	return result, nil
}

// columnToIndex 欄位字串轉為索引值
func columnToIndex(letter string) int {
	if columnChecker(letter) == false {
		panic("invalid column") // 正常狀況下應該不會跑出異常
	} // if

	result := 0

	for _, itor := range strings.ToLower(letter) {
		value := int(itor - 'a' + 1)
		result = result*26 + value // 英文字母26個字
	} // for

	return result
}

var columnChecker = regexp.MustCompile("^[a-zA-Z]+$").MatchString // 檢查欄位字串
