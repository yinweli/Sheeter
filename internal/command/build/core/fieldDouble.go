package core

import "strconv"

// FieldDouble 64位元浮點數
type FieldDouble struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldDouble) TypeExcel() string {
	return "double"
}

// TypeCpp 取得c++欄位類型
func (this *FieldDouble) TypeCpp() string {
	return "double"
}

// TypeCs 取得c#欄位類型
func (this *FieldDouble) TypeCs() string {
	return "double"
}

// TypeGo 取得go欄位類型
func (this *FieldDouble) TypeGo() string {
	return "float64"
}

// Show 是否顯示
func (this *FieldDouble) Show() bool {
	return true
}

// PrimaryKey 是否是主要索引
func (this *FieldDouble) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldDouble) Transform(input string) (result interface{}, err error) {
	return strconv.ParseFloat(input, 64)
}
