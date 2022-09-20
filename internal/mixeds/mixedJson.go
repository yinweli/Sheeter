package mixeds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
)

// Json json綜合工具
type Json struct {
	mixed *Mixed // 綜合工具
}

// FileJsonCsStruct 取得json-cs結構程式碼檔名路徑
func (this *Json) FileJsonCsStruct() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtCs,
		path:       internal.PathJsonCs,
	})
}

// FileJsonCsReader 取得json-cs讀取器程式碼檔名路徑
func (this *Json) FileJsonCsReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
		path:       internal.PathJsonCs,
	})
}

// FileJsonGoStruct 取得json-go結構程式碼檔名路徑
func (this *Json) FileJsonGoStruct() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtGo,
		path:       internal.PathJsonGo,
	})
}

// FileJsonGoReader 取得json-go讀取器程式碼檔名路徑
func (this *Json) FileJsonGoReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
		path:       internal.PathJsonGo,
	})
}

// FileJsonData 取得json資料檔名路徑
func (this *Json) FileJsonData() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJsonData,
		path:       internal.PathJsonData,
	})
}

// FileJsonDataCode 取得程式碼可用的json資料檔名路徑
func (this *Json) FileJsonDataCode() string {
	return filepath.ToSlash(this.FileJsonData()) // 因為要把路徑寫到程式碼中, 所以要改變分隔符號的方式
}
