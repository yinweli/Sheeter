package core

import "strconv"

// Double 64位元浮點數
type Double struct {
}

// TypeExcel 取得excel欄位類型
func (this *Double) TypeExcel() string {
	return "double"
}

// TypeCpp 取得c++欄位類型
func (this *Double) TypeCpp() string {
	return "double"
}

// TypeCs 取得c#欄位類型
func (this *Double) TypeCs() string {
	return "double"
}

// TypeGo 取得go欄位類型
func (this *Double) TypeGo() string {
	return "float64"
}

// Hide 是否隱藏
func (this *Double) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Double) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *Double) Transform(input string) (result interface{}, err error) {
	return strconv.ParseFloat(input, 64)
}
