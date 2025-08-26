package nameds

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/layouts"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Named 命名工具
type Named struct {
	Output    string              // 輸出路徑
	ExcelName string              // excel名稱
	SheetName string              // sheet名稱
	Primary   *layouts.LayoutData // 主索引布局資料
}

// AppName 取得程式名稱
func (this *Named) AppName() string {
	return sheeter.Application
}

// Namespace 取得命名空間名稱
func (this *Named) Namespace() string {
	return sheeter.Namespace
}

// StructName 取得結構名稱
func (this *Named) StructName() string {
	return combine(&param{
		excelUpper: true,
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
	})
}

// StructNote 取得結構說明
func (this *Named) StructNote() string {
	return fmt.Sprintf("%v#%v", this.ExcelName, this.SheetName)
}

// ReaderName 取得讀取器名稱
func (this *Named) ReaderName() string {
	return combine(&param{
		excelUpper: true,
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		last:       sheeter.Reader,
	})
}

// JsonName 取得json名稱
func (this *Named) JsonName() string {
	return combine(&param{
		excelName: this.ExcelName,
		sheetName: this.SheetName,
	})
}

// JsonExt 取得json副檔名
func (this *Named) JsonExt() string {
	return sheeter.ExtJson
}

// DataFile 取得資料檔名
func (this *Named) DataFile() string {
	return combine(&param{
		excelName: this.ExcelName,
		sheetName: this.SheetName,
		ext:       sheeter.ExtJson,
	})
}

// DataPath 取得資料路徑
func (this *Named) DataPath() string {
	return filepath.Join(this.Output, sheeter.PathJson, this.DataFile())
}

// ReaderPathCs 取得cs語言讀取程式碼路徑
func (this *Named) ReaderPathCs() string {
	return filepath.Join(this.Output, sheeter.PathCs, combine(&param{
		excelUpper: true,
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		last:       sheeter.Reader,
		ext:        sheeter.ExtCs,
	}))
}

// SheeterPathCs 取得cs語言表格程式碼路徑
func (this *Named) SheeterPathCs() string {
	return filepath.Join(this.Output, sheeter.PathCs, utils.FirstUpper(sheeter.Sheeter)+sheeter.ExtCs)
}

// HelperPathCs 取得cs語言工具程式碼路徑
func (this *Named) HelperPathCs() string {
	return filepath.Join(this.Output, sheeter.PathCs, utils.FirstUpper(sheeter.Helper)+sheeter.ExtCs)
}

// ReaderPathGo 取得go語言讀取程式碼路徑
func (this *Named) ReaderPathGo() string {
	return filepath.Join(this.Output, sheeter.PathGo, combine(&param{
		excelName: this.ExcelName,
		sheetName: this.SheetName,
		last:      sheeter.Reader,
		ext:       sheeter.ExtGo,
	}))
}

// SheeterPathGo 取得go語言表格程式碼路徑
func (this *Named) SheeterPathGo() string {
	return filepath.Join(this.Output, sheeter.PathGo, sheeter.Sheeter+sheeter.ExtGo)
}

// HelperPathGo 取得cs語言工具程式碼路徑
func (this *Named) HelperPathGo() string {
	return filepath.Join(this.Output, sheeter.PathGo, sheeter.Helper+sheeter.ExtGo)
}

// PrimaryCs 取得cs主索引類型字串
func (this *Named) PrimaryCs() string {
	return this.Primary.Field.ToTypeCs()
}

// PrimaryGo 取得go主索引類型字串
func (this *Named) PrimaryGo() string {
	return this.Primary.Field.ToTypeGo()
}

// FirstUpper 字串首字母大寫
func (this *Named) FirstUpper(input string) string {
	return utils.FirstUpper(input)
}

// FirstLower 字串首字母小寫
func (this *Named) FirstLower(input string) string {
	return utils.FirstLower(input)
}
