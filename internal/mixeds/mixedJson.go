package mixeds

import (
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
		path:       []string{internal.PathJson, internal.PathCs},
	})
}

// FileJsonCsReader 取得json-cs讀取器程式碼檔名路徑
func (this *Json) FileJsonCsReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
		path:       []string{internal.PathJson, internal.PathCs},
	})
}

// FileJsonGoStruct 取得json-go結構程式碼檔名路徑
func (this *Json) FileJsonGoStruct() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtGo,
		path:       []string{internal.PathJson, internal.PathGo},
	})
}

// FileJsonGoReader 取得json-go讀取器程式碼檔名路徑
func (this *Json) FileJsonGoReader() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
		path:       []string{internal.PathJson, internal.PathGo},
	})
}

// FileJsonDataName 取得json資料檔案名稱
func (this *Json) FileJsonDataName() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJsonData,
	})
}

// FileJsonDataPath 取得json資料檔名路徑
func (this *Json) FileJsonDataPath() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJsonData,
		path:       []string{internal.PathJson, internal.PathData},
	})
}
