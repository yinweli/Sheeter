package tasks

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/util"

	"github.com/xuri/excelize/v2"
)

// originalName 取得原始名稱
func (this *Task) originalName() string {
	return fmt.Sprintf("%s(%s)", this.element.Excel, this.element.Sheet)
}

// namespace 取得命名空間名稱
func (this *Task) namespace() string {
	return internal.Title
}

// structName 取得結構名稱
func (this *Task) structName() string {
	excelName := util.FirstUpper(this.excelName())
	sheetName := util.FirstUpper(this.element.Sheet)

	return excelName + sheetName
}

// readerName 取得讀取器名稱
func (this *Task) readerName() string {
	return this.structName() + "Reader"
}

// excelName 取得沒有副檔名的excel名稱
func (this *Task) excelName() string {
	return strings.TrimSuffix(this.element.Excel, filepath.Ext(this.element.Excel))
}

// sheetExists 表格是否存在
func (this *Task) sheetExists() bool {
	return this.excel.GetSheetIndex(this.element.Sheet) != -1
}

// getRows 取得表格行資料, line從1起算; 如果該行不存在, 回傳成功並取得最後一行物件
func (this *Task) getRows(line int) (rows *excelize.Rows, err error) {
	if line <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("row <= 0")
	} // if

	rows, err = this.excel.Rows(this.element.Sheet)

	if err != nil {
		return nil, fmt.Errorf("get rows failed: %w", err)
	} // if

	for l := 0; l < line; l++ {
		rows.Next()
	} // for

	return rows, nil
}

// getRowContent 取得表格行內容, line從1起算; 如果該行不存在, 回傳失敗
func (this *Task) getRowContent(line int) (cols []string, err error) {
	if line <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("row <= 0")
	} // if

	rows, err := this.excel.Rows(this.element.Sheet)

	if err != nil {
		return nil, fmt.Errorf("get rows failed: %w", err)
	} // if

	defer func() { _ = rows.Close() }()

	for l := 0; l < line; l++ {
		if rows.Next() == false {
			return nil, fmt.Errorf("row not found")
		} // if
	} // for

	cols, err = rows.Columns()

	if err != nil {
		return nil, fmt.Errorf("get columns failed: %w", err)
	} // if

	if cols == nil {
		cols = []string{} // 如果取得空行, 就回傳個空切片吧
	} // if

	return cols, nil
}
