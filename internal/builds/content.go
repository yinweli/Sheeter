package builds

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/builds/layouts"
	"github.com/yinweli/Sheeter/internal/util"
)

const pathJson = "json"              // 輸出路徑: json資料
const pathJsonSchema = "json-schema" // 輸出路徑: json架構
const pathJsonCs = "json-cs"         // 輸出路徑: json-cs
const pathJsonGo = "json-go"         // 輸出路徑: json-go
const fileJsonCsCode = "sheeter.cs"  // 輸出檔名: json-cs程式碼
const fileJsonCsReader = "reader.cs" // 輸出檔名: json-cs讀取器
const fileJsonGoCode = "sheeter.go"  // 輸出檔名: json-go程式碼
const fileJsonGoReader = "reader.go" // 輸出檔名: json-go讀取器
const extJson = "json"               // 副檔名: json

// Content 內容資料
type Content struct {
	Bom         bool             // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int              // 欄位行號(1為起始行)
	LineOfLayer int              // 階層行號(1為起始行)
	LineOfNote  int              // 註解行號(1為起始行)
	LineOfData  int              // 資料行號(1為起始行)
	Excel       string           // excel檔案路徑
	Sheet       string           // excel表單名稱
	excel       *excelize.File   // excel物件
	builder     *layouts.Builder // 布局建造器
}

// StructName 取得結構名稱
func (this *Content) StructName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
	})
}

// FileJson 取得json檔名路徑
func (this *Content) FileJson() string {
	name := this.combine(params{
		excelUpper: false,
		sheetUpper: true,
		ext:        extJson,
	})
	return filepath.Join(pathJson, name)
}

// FileJsonSchema 取得json架構檔名路徑
func (this *Content) FileJsonSchema() string {
	name := this.combine(params{
		excelUpper: false,
		sheetUpper: true,
		ext:        extJson,
	})
	return filepath.Join(pathJsonSchema, name)
}

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	sheetUpper bool   // sheet名稱是否要首字大寫
	ext        string // 副檔名
}

// combine 取得組合名稱
func (this *Content) combine(params params) string {
	excel := strings.TrimSuffix(filepath.Base(this.Excel), filepath.Ext(this.Excel))

	if params.excelUpper {
		excel = util.FirstUpper(excel)
	} else {
		excel = util.FirstLower(excel)
	} // if

	sheet := this.Sheet

	if params.sheetUpper {
		sheet = util.FirstUpper(sheet)
	} else {
		sheet = util.FirstLower(sheet)
	} // if

	ext := ""

	if params.ext != "" {
		ext = "." + params.ext
	} // if

	return excel + sheet + ext
}

// Close 關閉excel物件
func (this *Content) Close() {
	if this.excel != nil {
		_ = this.excel.Close()
	} // if
}

// GetRows 取得表格行資料, line從1起算; 如果該行不存在, 回傳成功並取得最後一行物件
func (this *Content) GetRows(line int) (rows *excelize.Rows, err error) {
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
func (this *Content) GetColumns(line int) (cols []string, err error) {
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

// Contents 內容列表
type Contents struct {
	Contents []*Content // 內容列表
	maxline  int        // 最大行數
	curline  int        // 當前行數
}

// AppName 取得程式名稱
func (this *Contents) AppName() string {
	return internal.Title
}

// Namespace 取得命名空間名稱
func (this *Contents) Namespace() string {
	return internal.Title
}

// PathJson 取得json路徑
func (this *Contents) PathJson() string {
	return pathJson
}

// PathJsonSchema 取得json架構路徑
func (this *Contents) PathJsonSchema() string {
	return pathJsonSchema
}

// PathJsonCs 取得json-cs路徑
func (this *Contents) PathJsonCs() string {
	return pathJsonCs
}

// PathJsonGo 取得json-go路徑
func (this *Contents) PathJsonGo() string {
	return pathJsonGo
}

// FileJsonCsCode 取得json-cs程式碼檔名路徑
func (this *Contents) FileJsonCsCode() string {
	return filepath.Join(pathJsonCs, fileJsonCsCode)
}

// FileJsonCsReader 取得json-cs讀取器檔名路徑
func (this *Contents) FileJsonCsReader() string {
	return filepath.Join(pathJsonCs, fileJsonCsReader)
}

// FileJsonGoCode 取得json-go程式碼檔名路徑
func (this *Contents) FileJsonGoCode() string {
	return filepath.Join(pathJsonGo, fileJsonGoCode)
}

// FileJsonGoReader 取得json-go讀取器檔名路徑
func (this *Contents) FileJsonGoReader() string {
	return filepath.Join(pathJsonGo, fileJsonGoReader)
}

// SetLine 設置行數
func (this *Contents) SetLine() string {
	this.maxline = len(this.Contents) - 1
	this.curline = 0
	return ""
}

// NewLine 換行輸出
func (this *Contents) NewLine() string {
	defer func() { this.curline++ }()

	if this.maxline > this.curline {
		return "\n"
	} // if

	return ""
}
