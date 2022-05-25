package core

import "strconv"

// Bool 布林值
type Bool struct {
}

// TypeExcel 取得excel欄位類型
func (this *Bool) TypeExcel() string {
	return "bool"
}

// TypeCpp 取得c++欄位類型
func (this *Bool) TypeCpp() string {
	return "bool"
}

// TypeCs 取得c#欄位類型
func (this *Bool) TypeCs() string {
	return "bool"
}

// TypeGo 取得go欄位類型
func (this *Bool) TypeGo() string {
	return "bool"
}

// Hide 是否隱藏
func (this *Bool) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Bool) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *Bool) Transform(input string) (result interface{}, err error) {
	return strconv.ParseBool(input)
}
