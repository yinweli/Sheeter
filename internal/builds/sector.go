package builds

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Sector 區段資料
type Sector struct {
	Global                         // 全域設定
	Element                        // 項目設定
	excel      *excelize.File      // excel物件
	layoutJson *layouts.LayoutJson // json布局器
	layoutType *layouts.LayoutType // 類型布局器
}

// StructName 取得結構名稱
func (this *Sector) StructName() string {
	name := this.combine(params{
		excelUpper: true,
		sheetUpper: true,
	})
	return name
}

// ReaderName 取得讀取器名稱
func (this *Sector) ReaderName() string {
	name := this.combine(params{
		excelUpper: true,
		sheetUpper: true,
		last:       internal.Reader,
	})
	return name
}

// FileJson 取得json檔名路徑
func (this *Sector) FileJson() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJson,
	})
	path := filepath.Join(internal.PathJson, name)
	return path
}

// FileJsonSchema 取得json架構檔名路徑
func (this *Sector) FileJsonSchema() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJson,
	})
	return filepath.Join(internal.PathJsonSchema, name)
}

// Close 關閉excel物件
func (this *Sector) Close() {
	if this.excel != nil {
		_ = this.excel.Close()
	} // if
}

// GetRows 取得表格行資料, line從1起算; 如果該行不存在, 回傳成功並取得最後一行物件
func (this *Sector) GetRows(line int) (rows *excelize.Rows, err error) {
	if line <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("get row failed, row <= 0")
	} // if

	rows, err = this.excel.Rows(this.Sheet)

	if err != nil {
		return nil, fmt.Errorf("get row failed: %w", err)
	} // if

	for l := 0; l < line; l++ {
		rows.Next()
	} // for

	return rows, nil
}

// GetColumns 取得表格行內容, line從1起算; 如果該行不存在, 回傳失敗
func (this *Sector) GetColumns(line int) (cols []string, err error) {
	if line <= 0 { // 注意! 最少要一次才能定位到第1行; 所以若line <= 0, 就表示錯誤
		return nil, fmt.Errorf("get columns failed, row <= 0")
	} // if

	rows, err := this.excel.Rows(this.Sheet)

	if err != nil {
		return nil, fmt.Errorf("get columns failed: %w", err)
	} // if

	defer func() { _ = rows.Close() }()

	for l := 0; l < line; l++ {
		if rows.Next() == false {
			return nil, fmt.Errorf("get columns failed, row not found")
		} // if
	} // for

	cols, err = rows.Columns()

	if err != nil {
		return nil, fmt.Errorf("get columns failed, invalid columns: %w", err)
	} // if

	if cols == nil {
		cols = []string{} // 如果取得空行, 就回傳個空切片吧
	} // if

	return cols, nil
}

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	sheetUpper bool   // sheet名稱是否要首字大寫
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
}

// combine 取得組合名稱
func (this *Sector) combine(params params) string {
	excel := utils.FileName(this.Excel)

	if params.excelUpper {
		excel = utils.FirstUpper(excel)
	} else {
		excel = utils.FirstLower(excel)
	} // if

	sheet := this.Sheet

	if params.sheetUpper {
		sheet = utils.FirstUpper(sheet)
	} else {
		sheet = utils.FirstLower(sheet)
	} // if

	items := []string{excel, sheet, params.last}

	if params.ext != "" {
		items = append(items, ".", params.ext)
	} // if

	return strings.Join(items, "")
}

// MergeSectorLayoutType 合併區段資料的類型布局器
func MergeSectorLayoutType(sectors []*Sector) (layoutType *layouts.LayoutType, err error) {
	layoutType = layouts.NewLayoutType()

	for _, itor := range sectors {
		if err := layoutType.Merge(itor.layoutType); err != nil {
			return nil, fmt.Errorf("merge sector layoutType failed: %w", err)
		} // if
	} // for

	return layoutType, nil
}
