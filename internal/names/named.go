package names

import (
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Named 命名工具
type Named struct {
	Excel string // excel檔案名稱
	Sheet string // excel表單名稱
}

// AppName 取得程式名稱
func (this *Named) AppName() string {
	return internal.AppName
}

// Namespace 取得命名空間名稱
func (this *Named) Namespace() string {
	return internal.AppName
}

// StructName 取得結構名稱
func (this *Named) StructName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
	})
}

// ReaderName 取得讀取器名稱
func (this *Named) ReaderName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
		last:       internal.Reader,
	})
}

// FileJson 取得json檔名路徑
func (this *Named) FileJson() string {
	return this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJson,
		path:       internal.PathJson,
	})
}

// FileJsonCode 取得程式碼可用的json檔名路徑
func (this *Named) FileJsonCode() string {
	return filepath.ToSlash(this.FileJson()) // 因為要把路徑寫到程式碼中, 所以要改變分隔符號的方式
}

// FileJsonSchema 取得json架構檔名路徑
func (this *Named) FileJsonSchema() string {
	return this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtJson,
		path:       internal.PathJsonSchema,
	})
}

// FileJsonCsStruct 取得json-cs程式碼檔名路徑
func (this *Named) FileJsonCsStruct() string {
	return this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtCs,
		path:       internal.PathJsonCs,
	})
}

// FileJsonCsReader 取得json-cs讀取器檔名路徑
func (this *Named) FileJsonCsReader() string {
	return this.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtCs,
		path:       internal.PathJsonCs,
	})
}

// FileJsonGoStruct 取得json-go程式碼檔名路徑
func (this *Named) FileJsonGoStruct() string {
	return this.combine(params{
		sheetUpper: true,
		ext:        internal.ExtGo,
		path:       internal.PathJsonGo,
	})
}

// FileJsonGoReader 取得json-go讀取器檔名路徑
func (this *Named) FileJsonGoReader() string {
	return this.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		ext:        internal.ExtGo,
		path:       internal.PathJsonGo,
	})
}

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	sheetUpper bool   // sheet名稱是否要首字大寫
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
	path       string // 路徑字串
}

// combine 取得組合名稱
func (this *Named) combine(params params) string {
	excel := utils.FileName(this.Excel)

	if params.excelUpper {
		excel = utils.FirstUpper(excel)
	} else {
		excel = utils.FirstLower(excel)
	} // if

	sheet := this.Sheet

	if params.sheetUpper {
		sheet = utils.FirstUpper(sheet)
	} else {
		sheet = utils.FirstLower(sheet)
	} // if

	name := excel + sheet + params.last

	if params.ext != "" {
		name = name + "." + params.ext
	} // if

	path := filepath.Join(params.path, name)
	return path
}
