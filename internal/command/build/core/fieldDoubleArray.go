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

// ToJsonValue 轉換為json值
func (this *FieldDoubleArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToFloatArray(input)
}
