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
		path:       internal.PathJson,
		ext:        internal.ExtJson,
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
		path:       internal.PathJsonSchema,
		ext:        internal.ExtJson,
	})
}

// FileJsonCsCode 取得json-cs程式碼檔名路徑
func (this *Named) FileJsonCsCode() string {
	return this.combine(params{
		sheetUpper: true,
		path:       internal.PathJsonCs,
		ext:        internal.ExtCs,
	})
}

// FileJsonCsReader 取得json-cs讀取器檔名路徑
func (this *Named) FileJsonCsReader() string {
	return this.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		path:       internal.PathJsonCs,
		ext:        internal.ExtCs,
	})
}

// FileJsonGoCode 取得json-go程式碼檔名路徑
func (this *Named) FileJsonGoCode() string {
	return this.combine(params{
		sheetUpper: true,
		path:       internal.PathJsonGo,
		ext:        internal.ExtGo,
	})
}

// FileJsonGoReader 取得json-go讀取器檔名路徑
func (this *Named) FileJsonGoReader() string {
	return this.combine(params{
		sheetUpper: true,
		last:       internal.Reader,
		path:       internal.PathJsonGo,
		ext:        internal.ExtGo,
	})
}

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	sheetUpper bool   // sheet名稱是否要首字大寫
	last       string // excel與sheet的結尾字串
	path       string // 路徑字串
	ext        string // 副檔名
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
