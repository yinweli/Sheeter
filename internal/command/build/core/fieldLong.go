package core

import (
	"strconv"
)

// FieldLong 64位元整數
type FieldLong struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldLong) TypeExcel() string {
	return "long"
}

// TypeCpp 取得c++欄位類型
func (this *FieldLong) TypeCpp() string {
	return "int64_t"
}

// TypeCs 取得c#欄位類型
func (this *FieldLong) TypeCs() string {
	return "long"
}

// TypeGo 取得go欄位類型
func (this *FieldLong) TypeGo() string {
	return "int64"
}

// IsShow 是否顯示
func (this *FieldLong) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldLong) IsPkey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldLong) Transform(input string) (result interface{}, err error) {
	return strconv.ParseInt(input, 10, 64)
}
