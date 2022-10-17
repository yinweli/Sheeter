package mixeds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
)

// Json json綜合工具
type Json struct {
	mixed *Mixed // 綜合工具
}

// JsonDataName 取得json資料名稱
func (this *Json) JsonDataName() string {
	return this.mixed.combine(params{
		sheetUpper: true,
	})
}

// JsonDataExt 取得json資料副檔名
func (this *Json) JsonDataExt() string {
	return internal.JsonDataExt
}

// JsonDataFile 取得json資料檔名
func (this *Json) JsonDataFile() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.JsonDataExt,
	})
}

// JsonDataPath 取得json資料路徑
func (this *Json) JsonDataPath() string {
	return filepath.Join(internal.JsonPath, internal.DataPath, this.JsonDataFile())
}

// JsonStructCsPath 取得json結構cs程式碼路徑
func (this *Json) JsonStructCsPath() string {
	return filepath.Join(internal.JsonPath, internal.CsPath, this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.CsExt,
	}))
}

// JsonReaderCsPath 取得json讀取器cs程式碼路徑
func (this *Json) JsonReaderCsPath() string {
	return filepath.Join(internal.JsonPath, internal.CsPath, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.CsExt,
	}))
}

// JsonDepotCsPath 取得json倉庫cs程式碼路徑
func (this *Json) JsonDepotCsPath() string {
	return filepath.Join(internal.JsonPath, internal.CsPath, internal.Depot+"."+internal.CsExt)
}

// JsonStructGoPath 取得json-go結構程式碼路徑
func (this *Json) JsonStructGoPath() string {
	return filepath.Join(internal.JsonPath, internal.GoPath, this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.GoExt,
	}))
}

// JsonReaderGoPath 取得json-go讀取器程式碼檔名路徑
func (this *Json) JsonReaderGoPath() string {
	return filepath.Join(internal.JsonPath, internal.GoPath, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.GoExt,
	}))
}

// JsonDepotGoPath 取得json-go倉庫程式碼路徑
func (this *Json) JsonDepotGoPath() string {
	return filepath.Join(internal.JsonPath, internal.GoPath, internal.Depot+"."+internal.GoExt)
}
