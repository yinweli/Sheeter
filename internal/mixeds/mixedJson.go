package mixeds

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
)

// Json json綜合工具
type Json struct {
	mixed *Mixed // 綜合工具
}

// FileJsonData 取得json資料檔名
func (this *Json) FileJsonData() string {
	return this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJsonData,
	})
}

// PathJsonData 取得json資料路徑
func (this *Json) PathJsonData() string {
	return filepath.Join(internal.PathJson, internal.PathData, this.FileJsonData())
}

// PathJsonCsStruct 取得json-cs結構程式碼路徑
func (this *Json) PathJsonCsStruct() string {
	return filepath.Join(internal.PathJson, internal.PathCs, this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtCs,
	}))
}

// PathJsonCsReader 取得json-cs讀取器程式碼路徑
func (this *Json) PathJsonCsReader() string {
	return filepath.Join(internal.PathJson, internal.PathCs, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
	}))
}

// PathJsonGoStruct 取得json-go結構程式碼路徑
func (this *Json) PathJsonGoStruct() string {
	return filepath.Join(internal.PathJson, internal.PathGo, this.mixed.combine(params{
		sheetUpper: true,
		ext:        internal.ExtGo,
	}))
}

// PathJsonGoReader 取得json-go讀取器程式碼檔名路徑
func (this *Json) PathJsonGoReader() string {
	return filepath.Join(internal.PathJson, internal.PathGo, this.mixed.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
	}))
}
