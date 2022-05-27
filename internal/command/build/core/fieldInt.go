package core

import (
	"strconv"
)

// FieldInt 32位元整數
type FieldInt struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldInt) TypeExcel() string {
	return "int"
}

// TypeCpp 取得c++欄位類型
func (this *FieldInt) TypeCpp() string {
	return "int32_t"
}

// TypeCs 取得c#欄位類型
func (this *FieldInt) TypeCs() string {
	return "int"
}

// TypeGo 取得go欄位類型
func (this *FieldInt) TypeGo() string {
	return "int32"
}

// Hide 是否隱藏
func (this *FieldInt) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldInt) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldInt) Transform(input string) (result interface{}, err error) {
	return strconv.ParseInt(input, 10, 32)
}
