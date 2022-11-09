package nameds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/sheeter"
)

// Enum enum命名工具
type Enum struct {
	ExcelName string // excel名稱
	SheetName string // sheet名稱
}

// EnumCsPath 取得enum-cs路徑
func (this *Enum) EnumCsPath() string {
	return filepath.Join(sheeter.EnumPath, sheeter.CsPath)
}

// EnumGoPath 取得enum-go路徑
func (this *Enum) EnumGoPath() string {
	return filepath.Join(sheeter.EnumPath, sheeter.GoPath)
}

// EnumSchemaPath 取得enum-schema路徑
func (this *Enum) EnumSchemaPath() string {
	return filepath.Join(sheeter.EnumPath, sheeter.SchemaPath)
}

// EnumName 取得enum架構檔名
func (this *Enum) EnumName() string {
	return combine(&params{
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		sheetUpper: true,
		ext:        sheeter.EnumSchemaExt,
	})
}

// EnumPath 取得enum架構路徑
func (this *Enum) EnumPath() string {
	return filepath.Join(this.EnumSchemaPath(), this.EnumName())
}
