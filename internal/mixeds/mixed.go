package mixeds

import (
	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// NewMixed 建立綜合工具
func NewMixed(excel, sheet string) *Mixed {
	mixed := &Mixed{
		excelName: excel,
		sheetName: sheet,
	}
	mixed.Json.mixed = mixed
	mixed.Proto.mixed = mixed
	return mixed
}

// Mixed 綜合工具, 用在程式中以及模板中
type Mixed struct {
	Json             // json綜合工具
	Proto            // proto綜合工具
	Field            // 欄位綜合工具
	Math             // 數學綜合工具
	String           // 字串綜合工具
	excelName string // excel檔案名稱
	sheetName string // excel表單名稱
}

// AppName 取得程式名稱
func (this *Mixed) AppName() string {
	return internal.AppName
}

// JsonNamespace 取得json命名空間名稱
func (this *Mixed) JsonNamespace(simpleNamespace bool) string {
	if simpleNamespace {
		return internal.AppName
	} else {
		return internal.JsonNamespace
	} // if
}

// ProtoNamespace 取得proto命名空間名稱
func (this *Mixed) ProtoNamespace(simpleNamespace bool) string {
	if simpleNamespace {
		return internal.AppName
	} else {
		return internal.ProtoNamespace
	} // if
}

// StructName 取得結構名稱
func (this *Mixed) StructName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
	})
}

// ReaderName 取得讀取器名稱
func (this *Mixed) ReaderName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
		last:       internal.Reader,
	})
}

// StorerName 取得儲存器名稱
func (this *Mixed) StorerName() string {
	return this.combine(params{
		excelUpper: true,
		sheetUpper: true,
		last:       internal.Storer,
	})
}

// StorerDatas 取得儲存器資料名稱
func (this *Mixed) StorerDatas() string {
	return internal.StorerDatas
}

// StorerMessage 取得儲存器proto message名稱
func (this *Mixed) StorerMessage(simpleNamespace bool) string {
	return this.ProtoNamespace(simpleNamespace) + "." + this.StorerName()
}

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	sheetUpper bool   // sheet名稱是否要首字大寫
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
}

// combine 取得組合名稱
func (this *Mixed) combine(params params) string {
	excel := utils.FileName(this.excelName)

	if params.excelUpper {
		excel = utils.FirstUpper(excel)
	} else {
		excel = utils.FirstLower(excel)
	} // if

	sheet := this.sheetName

	if params.sheetUpper {
		sheet = utils.FirstUpper(sheet)
	} else {
		sheet = utils.FirstLower(sheet)
	} // if

	name := excel + sheet + params.last

	if params.ext != "" {
		name = name + "." + params.ext
	} // if

	return name
}
