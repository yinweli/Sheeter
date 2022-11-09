package nameds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Proto proto命名工具
type Proto struct {
	ExcelName string // excel名稱
	SheetName string // sheet名稱
}

// ProtoCsPath 取得proto-cs路徑
func (this *Proto) ProtoCsPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.CsPath)
}

// ProtoGoPath 取得proto-go路徑
func (this *Proto) ProtoGoPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.GoPath)
}

// ProtoSchemaPath 取得proto-schema路徑
func (this *Proto) ProtoSchemaPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.SchemaPath)
}

// ProtoName 取得proto架構檔名
func (this *Proto) ProtoName() string {
	return combine(&params{
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		sheetUpper: true,
		ext:        sheeter.ProtoSchemaExt,
	})
}

// ProtoPath 取得proto架構路徑
func (this *Proto) ProtoPath() string {
	return filepath.Join(this.ProtoSchemaPath(), this.ProtoName())
}

// ProtoDataName 取得proto資料名稱
func (this *Proto) ProtoDataName() string {
	return combine(&params{
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		sheetUpper: true,
	})
}

// ProtoDataExt 取得proto資料副檔名
func (this *Proto) ProtoDataExt() string {
	return sheeter.ProtoDataExt
}

// ProtoDataFile 取得proto資料檔名
func (this *Proto) ProtoDataFile() string {
	return combine(&params{
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		sheetUpper: true,
		ext:        sheeter.ProtoDataExt,
	})
}

// ProtoDataPath 取得proto資料路徑
func (this *Proto) ProtoDataPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.DataPath, this.ProtoDataFile())
}

// ProtoReaderCsPath 取得proto讀取器cs程式碼路徑
func (this *Proto) ProtoReaderCsPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.CsPath, combine(&params{
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		excelUpper: true, // cs程式碼一律大寫開頭
		sheetUpper: true,
		last:       sheeter.Reader,
		ext:        sheeter.CsExt,
	}))
}

// ProtoDepotCsPath 取得proto倉庫cs程式碼路徑
func (this *Proto) ProtoDepotCsPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.CsPath, utils.FirstUpper(sheeter.Depot)+sheeter.CsExt) // cs程式碼一律大寫開頭
}

// ProtoReaderGoPath 取得proto-go讀取器程式碼路徑
func (this *Proto) ProtoReaderGoPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.GoPath, combine(&params{
		excelName:  this.ExcelName,
		sheetName:  this.SheetName,
		sheetUpper: true,
		last:       sheeter.Reader,
		ext:        sheeter.GoExt,
	}))
}

// ProtoDepotGoPath 取得proto-go倉庫程式碼路徑
func (this *Proto) ProtoDepotGoPath() string {
	return filepath.Join(sheeter.ProtoPath, sheeter.GoPath, sheeter.Depot+sheeter.GoExt)
}

// ProtoDepend 取得proto依賴檔名
func (this *Proto) ProtoDepend(name string) string {
	// proto依賴檔名必須跟已建立的proto檔名相符
	// 因為proto檔名是小寫駝峰, 所以這裡也必須是小寫駝峰

	return utils.FirstLower(name) + sheeter.ProtoSchemaExt
}
