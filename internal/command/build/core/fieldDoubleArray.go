package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldDoubleArray 64位元浮點數陣列
type FieldDoubleArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldDoubleArray) TypeExcel() string {
	return "doubleArray"
}

// TypeCpp 取得c++欄位類型
func (this *FieldDoubleArray) TypeCpp() string {
	return "std::vector<double>"
}

// TypeCs 取得c#欄位類型
func (this *FieldDoubleArray) TypeCs() string {
	return "List<double>"
}

// TypeGo 取得go欄位類型
func (this *FieldDoubleArray) TypeGo() string {
	return "[]float64"
}

// IsShow 是否顯示
func (this *FieldDoubleArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldDoubleArray) IsPkey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldDoubleArray) Transform(input string) (result interface{}, err error) {
	return util.StrToFloatArray(input)
}
