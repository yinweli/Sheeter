package core

import (
	"strconv"

	"Sheeter/internal"
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

// Hide 是否隱藏
func (this *FieldLong) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldLong) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldLong) Transform(input string) (result interface{}, err error) {
	return strconv.ParseInt(input, internal.Decimal, 64)
}
