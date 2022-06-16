package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldLongArray 64位元整數陣列
type FieldLongArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldLongArray) TypeExcel() string {
	return "longArray"
}

// TypeCs 取得c#欄位類型
func (this *FieldLongArray) TypeCs() string {
	return "List<long>"
}

// TypeGo 取得go欄位類型
func (this *FieldLongArray) TypeGo() string {
	return "[]int64"
}

// IsShow 是否顯示
func (this *FieldLongArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldLongArray) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldLongArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToIntArray(input)
}
