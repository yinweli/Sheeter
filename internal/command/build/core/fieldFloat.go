package core

import "strconv"

// Float 32位元浮點數
type Float struct {
}

// TypeExcel 取得excel欄位類型
func (this *Float) TypeExcel() string {
	return "float"
}

// TypeCpp 取得c++欄位類型
func (this *Float) TypeCpp() string {
	return "float"
}

// TypeCs 取得c#欄位類型
func (this *Float) TypeCs() string {
	return "float"
}

// TypeGo 取得go欄位類型
func (this *Float) TypeGo() string {
	return "float32"
}

// Hide 是否隱藏
func (this *Float) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Float) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *Float) Transform(input string) (result interface{}, err error) {
	return strconv.ParseFloat(input, 32)
}
