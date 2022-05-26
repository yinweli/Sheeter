package core

import "strconv"

// FieldFloat 32位元浮點數
type FieldFloat struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldFloat) TypeExcel() string {
	return "float"
}

// TypeCpp 取得c++欄位類型
func (this *FieldFloat) TypeCpp() string {
	return "float"
}

// TypeCs 取得c#欄位類型
func (this *FieldFloat) TypeCs() string {
	return "float"
}

// TypeGo 取得go欄位類型
func (this *FieldFloat) TypeGo() string {
	return "float32"
}

// Hide 是否隱藏
func (this *FieldFloat) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldFloat) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldFloat) Transform(input string) (result interface{}, err error) {
	return strconv.ParseFloat(input, 32)
}
