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

// JsonCsStructPath 取得json-cs結構程式碼路徑
func (this *Json) JsonCsStructPath() string {
	return filepath.Join(internal.JsonPath, internal.CsPath, this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.CsExt,
	}))
}

// JsonCsReaderPath 取得json-cs讀取器程式碼路徑
func (this *Json) JsonCsReaderPath() string {
	return filepath.Join(internal.JsonPath, internal.CsPath, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.CsExt,
	}))
}

// JsonCsDepotPath 取得json-cs倉庫程式碼路徑
func (this *Json) JsonCsDepotPath() string {
	return filepath.Join(internal.JsonPath, internal.CsPath, internal.Depot+"."+internal.CsExt)
}

// JsonGoStructPath 取得json-go結構程式碼路徑
func (this *Json) JsonGoStructPath() string {
	return filepath.Join(internal.JsonPath, internal.GoPath, this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.GoExt,
	}))
}

// JsonGoReaderPath 取得json-go讀取器程式碼檔名路徑
func (this *Json) JsonGoReaderPath() string {
	return filepath.Join(internal.JsonPath, internal.GoPath, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.GoExt,
	}))
}

// JsonGoDepotPath 取得json-go倉庫程式碼路徑
func (this *Json) JsonGoDepotPath() string {
	return filepath.Join(internal.JsonPath, internal.GoPath, internal.Depot+"."+internal.GoExt)
}
