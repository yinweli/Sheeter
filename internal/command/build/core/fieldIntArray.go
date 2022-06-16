package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldIntArray 32位元整數陣列
type FieldIntArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldIntArray) TypeExcel() string {
	return "intArray"
}

// TypeCs 取得c#欄位類型
func (this *FieldIntArray) TypeCs() string {
	return "List<int>"
}

// TypeGo 取得go欄位類型
func (this *FieldIntArray) TypeGo() string {
	return "[]int32"
}

// IsShow 是否顯示
func (this *FieldIntArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldIntArray) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldIntArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToIntArray(input)
}
