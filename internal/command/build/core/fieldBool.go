package core

import "strconv"

// FieldBool 布林值
type FieldBool struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldBool) TypeExcel() string {
	return "bool"
}

// TypeCpp 取得c++欄位類型
func (this *FieldBool) TypeCpp() string {
	return "bool"
}

// TypeCs 取得c#欄位類型
func (this *FieldBool) TypeCs() string {
	return "bool"
}

// TypeGo 取得go欄位類型
func (this *FieldBool) TypeGo() string {
	return "bool"
}

// Show 是否顯示
func (this *FieldBool) Show() bool {
	return true
}

// PrimaryKey 是否是主要索引
func (this *FieldBool) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldBool) Transform(input string) (result interface{}, err error) {
	return strconv.ParseBool(input)
}
