package nameds

import (
	"strconv"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Named 命名工具
type Named struct {
	ExcelName string // excel名稱
	SheetName string // sheet名稱
}

// AppName 取得程式名稱
func (this *Named) AppName() string {
	return internal.AppName
}

// JsonNamespace 取得json命名空間名稱
func (this *Named) JsonNamespace(simpleNamespace bool) string {
	if simpleNamespace {
		return internal.AppName
	} else {
		return internal.JsonNamespace
	} // if
}

// ProtoNamespace 取得proto命名空間名稱
func (this *Named) ProtoNamespace(simpleNamespace bool) string {
	if simpleNamespace {
		return internal.AppName
	} else {
		return internal.ProtoNamespace
	} // if
}

// EnumNamespace 取得enum命名空間名稱
func (this *Named) EnumNamespace(simpleNamespace bool) string {
	if simpleNamespace {
		return internal.AppName
	} else {
		return internal.EnumNamespace
	} // if
}

// StructName 取得結構名稱
func (this *Named) StructName() string {
	return combine(&params{
		excelName:  this.ExcelName,
		excelUpper: true,
		sheetName:  this.SheetName,
		sheetUpper: true,
	})
}

// ReaderName 取得讀取器名稱
func (this *Named) ReaderName() string {
	return combine(&params{
		excelName:  this.ExcelName,
		excelUpper: true,
		sheetName:  this.SheetName,
		sheetUpper: true,
		last:       internal.Reader,
	})
}

// StorerName 取得儲存器名稱
func (this *Named) StorerName() string {
	return combine(&params{
		excelName:  this.ExcelName,
		excelUpper: true,
		sheetName:  this.SheetName,
		sheetUpper: true,
		last:       internal.Storer,
	})
}

// StorerDatas 取得儲存器資料名稱
func (this *Named) StorerDatas() string {
	return internal.StorerDatas
}

// StorerMessage 取得儲存器proto message名稱
func (this *Named) StorerMessage(simpleNamespace bool) string {
	return this.ProtoNamespace(simpleNamespace) + "." + this.StorerName()
}

// FirstUpper 字串首字母大寫
func (this *Named) FirstUpper(input string) string {
	return utils.FirstUpper(input)
}

// FirstLower 字串首字母小寫
func (this *Named) FirstLower(input string) string {
	return utils.FirstLower(input)
}

// Add 加法
func (this *Named) Add(l, r int) string {
	return strconv.Itoa(l + r)
}

// Sub 減法
func (this *Named) Sub(l, r int) string {
	return strconv.Itoa(l - r)
}

// Mul 乘法
func (this *Named) Mul(l, r int) string {
	return strconv.Itoa(l * r)
}

// Div 除法
func (this *Named) Div(l, r int) string {
	return strconv.Itoa(l / r)
}
