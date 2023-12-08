package nameds

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// Named 命名工具
type Named struct {
	Output    string // 輸出路徑
	ExcelName string // excel名稱
	SheetName string // sheet名稱
}

// AppName 取得程式名稱
func (this *Named) AppName() string {
	return sheeter.AppName
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
	return fmt.Sprintf("%v # %v", this.ExcelName, this.SheetName)
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
	return sheeter.JsonExt
}

// DataFile 取得資料檔名
func (this *Named) DataFile() string {
	return combine(&param{
		excelName: this.ExcelName,
		sheetName: this.SheetName,
		ext:       sheeter.JsonExt,
	})
}

// DataPath 取得資料路徑
func (this *Named) DataPath() string {
	return filepath.Join(this.Output, sheeter.JsonPath, this.DataFile())
}

// ReaderPathCs 取得cs讀取器程式碼路徑
func (this *Named) ReaderPathCs() string {
	return filepath.Join(this.Output, sheeter.CsPath, combine(&param{
		excelUpper: true, // cs程式碼一律大寫開頭
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		last:       sheeter.Reader,
		ext:        sheeter.CsExt,
	}))
}

// SheeterPathCs 取得cs表格器程式碼路徑
func (this *Named) SheeterPathCs() string {
	return filepath.Join(this.Output, sheeter.CsPath, utils.FirstUpper(sheeter.Sheeter)+sheeter.CsExt) // cs程式碼一律大寫開頭
}

// ReaderPathGo 取得go讀取器程式碼路徑
func (this *Named) ReaderPathGo() string {
	return filepath.Join(this.Output, sheeter.GoPath, combine(&param{
		excelName: this.ExcelName,
		sheetName: this.SheetName,
		last:      sheeter.Reader,
		ext:       sheeter.GoExt,
	}))
}

// SheeterPathGo 取得go表格器程式碼路徑
func (this *Named) SheeterPathGo() string {
	return filepath.Join(this.Output, sheeter.GoPath, sheeter.Sheeter+sheeter.GoExt)
}

// FirstUpper 字串首字母大寫
func (this *Named) FirstUpper(input string) string {
	return utils.FirstUpper(input)
}

// FirstLower 字串首字母小寫
func (this *Named) FirstLower(input string) string {
	return utils.FirstLower(input)
}
