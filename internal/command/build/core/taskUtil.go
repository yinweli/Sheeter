package core

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
	"github.com/yinweli/Sheeter/internal/util"
)

// originalName 取得原始名稱
func (this *Task) originalName() string {
	return fmt.Sprintf("%s(%s)", this.element.Excel, this.element.Sheet)
}

// structName 取得結構名稱
func (this *Task) structName() string {
	excelName := util.FirstUpper(this.excelName())
	sheetName := util.FirstUpper(this.element.Sheet)

	return excelName + sheetName
}

// excelName 取得沒有副檔名的excel名稱
func (this *Task) excelName() string {
	return strings.TrimSuffix(this.element.Excel, filepath.Ext(this.element.Excel))
}

// sheetExists 表格是否存在
func (this *Task) sheetExists() bool {
	return this.excel.GetSheetIndex(this.element.Sheet) != -1
}

// getRows 取得表格行資料, line從1起算
func (this *Task) getRows(line int) *excelize.Rows {
	if line <= 0 {
		return nil
	} // if

	rows, err := this.excel.Rows(this.element.Sheet)

	if err != nil {
		return nil
	} // if

	for l := 0; l < line; l++ {
		if rows.Next() == false { // 注意! 最少要一次才能定位到第1行; 所以若line=0, 則取不到資料
			_ = rows.Close()
			return nil
		} // if
	} // for

	return rows
}

// getRowContent 取得表格行內容, line從1起算
func (this *Task) getRowContent(line int) []string {
	rows := this.getRows(line)

	if rows == nil {
		return nil
	} // if

	defer func() { _ = rows.Close() }()
	cols, err := rows.Columns()

	if err != nil {
		return nil
	} // if

	if cols == nil {
		cols = []string{} // 如果取得空行, 就回傳個空切片吧
	} // if

	return cols
}
