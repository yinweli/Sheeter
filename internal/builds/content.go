package builds

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/vbauerster/mpb/v7"
	"github.com/xuri/excelize/v2"
	"github.com/yinweli/Sheeter/internal"

	"github.com/yinweli/Sheeter/internal/builds/layouts"
	"github.com/yinweli/Sheeter/internal/util"
)

const pathSchema = "schema"  // 輸出路徑: json架構
const pathJson = "json"      // 輸出路徑: json
const pathJsonCs = "json-cs" // 輸出路徑: json-cs
const pathJsonGo = "json-go" // 輸出路徑: json-go
const lastReader = "Reader"  // 結尾名: 讀取器
const extSchema = "schema"   // 副檔名: json架構
const extJson = "json"       // 副檔名: json
const extCs = "cs"           // 副檔名: cs
const extGo = "go"           // 副檔名: go

// Content 內容資料
type Content struct {
	Bom         bool             // 輸出的檔案是否使用順序標記(BOM)
	LineOfField int              // 欄位行號(1為起始行)
	LineOfLayer int              // 階層行號(1為起始行)
	LineOfNote  int              // 註解行號(1為起始行)
	LineOfData  int              // 資料行號(1為起始行)
	Excel       string           // excel檔案路徑
	Sheet       string           // excel表單名稱
	Progress    *mpb.Progress    // 進度條產生器
	excel       *excelize.File   // excel物件
	builder     *layouts.Builder // 布局建造器
}

// Check 檢查工作
func (this *Content) Check() error {
	if this.LineOfField <= 0 {
		return fmt.Errorf("content failed, lineOfField <= 0")
	} // if

	if this.LineOfLayer <= 0 {
		return fmt.Errorf("content failed, lineOfLayer <= 0")
	} // if

	if this.LineOfNote <= 0 {
		return fmt.Errorf("content failed, lineOfNote <= 0")
	} // if

	if this.LineOfData <= 0 {
		return fmt.Errorf("content failed, lineOfData <= 0")
	} // if

	if this.LineOfField >= this.LineOfData {
		return fmt.Errorf("content failed, lineOfField(%d) >= lineOfData(%d)", this.LineOfField, this.LineOfData)
	} // if

	if this.LineOfLayer >= this.LineOfData {
		return fmt.Errorf("content failed, lineOfLayer(%d) >= lineOfData(%d)", this.LineOfLayer, this.LineOfData)
	} // if

	if this.LineOfNote >= this.LineOfData {
		return fmt.Errorf("content failed, lineOfNote(%d) >= lineOfData(%d)", this.LineOfNote, this.LineOfData)
	} // if

	if this.Excel == "" {
		return fmt.Errorf("content failed, excel empty")
	} // if

	if this.Sheet == "" {
		return fmt.Errorf("content failed, sheet empty")
	} // if

	if this.Progress == nil {
		return fmt.Errorf("content failed, progress nil")
	} // if

	return nil
}

// ShowName 顯示名稱
func (this *Content) ShowName() string {
	name := this.combine(params{
		middle: "#",
	})
	return name
}

// SchemaPath 取得json架構檔名路徑
func (this *Content) SchemaPath() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        extSchema,
	})
	return filepath.Join(pathSchema, name)
}

// JsonPath 取得json檔名路徑
func (this *Content) JsonPath() string {
	name := this.combine(params{
		sheetUpper: true,
		ext:        extJson,
	})
	return filepath.Join(pathJson, name)
}

// JsonCsPath 取得json-cs檔名路徑
func (this *Content) JsonCsPath() string {
	path := this.combine(params{
		sheetUpper: true,
	})
	name := this.combine(params{
		sheetUpper: true,
		ext:        extCs,
	})
	return filepath.Join(pathJsonCs, path, name)
}

// JsonCsReaderPath 取得json-cs讀取器檔名路徑
func (this *Content) JsonCsReaderPath() string {
	path := this.combine(params{
		sheetUpper: true,
	})
	name := this.combine(params{
		sheetUpper: true,
		last:       lastReader,
		ext:        extCs,
	})
	return filepath.Join(pathJsonCs, path, name)
}

// JsonGoPath 取得json-go檔名路徑
func (this *Content) JsonGoPath() string {
	path := this.combine(params{
		sheetUpper: true,
	})
	name := this.combine(params{
		sheetUpper: true,
		ext:        extGo,
	})
	return filepath.Join(pathJsonGo, path, name)
}

// JsonGoReaderPath 取得json-go讀取器檔名路徑
func (this *Content) JsonGoReaderPath() string {
	path := this.combine(params{
		sheetUpper: true,
	})
	name := this.combine(params{
		sheetUpper: true,
		last:       lastReader,
		ext:        extGo,
	})
	return filepath.Join(pathJsonGo, path, name)
}

// AppName 取得程式名稱
func (this *Content) AppName() string {
	return internal.Title
}

// Namespace 取得命名空間名稱
func (this *Content) Namespace() string {
	return this.combine(params{})
}

// StructName 取得結構名稱
func (this *Content) StructName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
	})
}

// ReaderName 取得讀取器名稱
func (this *Content) ReaderName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
		last:       lastReader,
	})
}

// getRows 取得表格行資料, line從1起算; 如果該行不存在, 回傳成功並取得最後一行物件
func (this *Content) getRows(line int) (rows *excelize.Rows, err error) {
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

// getColumns 取得表格行內容, line從1起算; 如果該行不存在, 回傳失敗
func (this *Content) getColumns(line int) (cols []string, err error) {
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

// close 關閉excel物件
func (this *Content) close() {
	if this.excel != nil {
		_ = this.excel.Close()
	} // if
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

	return excel + params.middle + sheet + params.last + ext
}

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	sheetUpper bool   // sheet名稱是否要首字大寫
	middle     string // excel與sheet的中間字串
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
}
