package core

import (
	"strconv"

	"Sheeter/internal"
)

// Int 32位元整數
type Int struct {
}

// TypeExcel 取得excel欄位類型
func (this *Int) TypeExcel() string {
	return "int"
}

// TypeCpp 取得c++欄位類型
func (this *Int) TypeCpp() string {
	return "int32_t"
}

// TypeCs 取得c#欄位類型
func (this *Int) TypeCs() string {
	return "int"
}

// TypeGo 取得go欄位類型
func (this *Int) TypeGo() string {
	return "int32"
}

// Hide 是否隱藏
func (this *Int) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Int) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *Int) Transform(input string) (result interface{}, err error) {
	return strconv.ParseInt(input, internal.Decimal, 32)
}
