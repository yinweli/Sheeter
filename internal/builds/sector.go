package builds

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/names"
)

// Sector 區段資料
type Sector struct {
	Global                         // 全域設定
	Element                        // 項目設定
	named      *names.Named        // 命名工具
	excel      *excelize.File      // excel物件
	layoutJson *layouts.LayoutJson // json布局器
	layoutType *layouts.LayoutType // 類型布局器
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
