package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldFloatArray 32位元浮點數陣列
type FieldFloatArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldFloatArray) TypeExcel() string {
	return "floatArray"
}

// TypeCs 取得c#欄位類型
func (this *FieldFloatArray) TypeCs() string {
	return "List<float>"
}

// TypeGo 取得go欄位類型
func (this *FieldFloatArray) TypeGo() string {
	return "[]float32"
}

// IsShow 是否顯示
func (this *FieldFloatArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldFloatArray) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldFloatArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToFloatArray(input)
}
