package core

import "Sheeter/internal/util"

// DoubleArray 64位元浮點數陣列
type DoubleArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *DoubleArray) TypeExcel() string {
	return "doubleArray"
}

// TypeCpp 取得c++欄位類型
func (this *DoubleArray) TypeCpp() string {
	return "std::vector<double>"
}

// TypeCs 取得c#欄位類型
func (this *DoubleArray) TypeCs() string {
	return "List<double>"
}

// TypeGo 取得go欄位類型
func (this *DoubleArray) TypeGo() string {
	return "[]float64"
}

// Hide 是否隱藏
func (this *DoubleArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *DoubleArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *DoubleArray) Transform(input string) (result interface{}, err error) {
	return util.StringToFloat64Array(input)
}
