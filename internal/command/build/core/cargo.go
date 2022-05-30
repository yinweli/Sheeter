package core

import (
	"fmt"
	"path/filepath"
	"strings"

	"Sheeter/internal/util"

	"github.com/schollz/progressbar/v3"
)

// Cargo 工作箱
type Cargo struct {
	Progress *progressbar.ProgressBar // 進度條
	Global   *Global                  // 全域設定
	Element  *Element                 // 項目設定
	Columns  []*Column                // 行資料列表
}

// JsonFileName 取得json檔案名稱
func (this *Cargo) JsonFileName() string {
	return this.fileName("json")
}

// CppFileName 取得c++檔案名稱
func (this *Cargo) CppFileName() string {
	return this.fileName("hpp")
}

// CsFileName 取得c#檔案名稱
func (this *Cargo) CsFileName() string {
	return this.fileName("cs")
}

// GoFileName 取得go檔案名稱
func (this *Cargo) GoFileName() string {
	return this.fileName("go")
}

// StructName 取得結構名稱
func (this *Cargo) StructName() string {
	excelName := util.FirstUpper(this.excelName())
	sheetName := util.FirstUpper(this.Element.Sheet)

	return excelName + sheetName
}

// fileName 取得輸出檔案名稱
func (this *Cargo) fileName(ext string) string {
	excelName := util.FirstLower(this.excelName())
	sheetName := util.FirstUpper(this.Element.Sheet)

	return fmt.Sprintf("%s%s.%s", excelName, sheetName, ext)
}

// outputExcelName 取得沒有副檔名的excel名稱
func (this *Cargo) excelName() string {
	return strings.TrimSuffix(this.Element.Excel, filepath.Ext(this.Element.Excel))
}

// Column 行資料
type Column struct {
	Note  string   // 欄位註解
	Name  string   // 欄位名稱
	Field Field    // 欄位類型
	Datas []string // 資料列表
}

// MemberName 取得成員名稱
func (this *Column) MemberName() string {
	return util.FirstUpper(this.Name)
}
