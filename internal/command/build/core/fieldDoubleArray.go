package core

import "Sheeter/internal/util"

// FieldDoubleArray 64位元浮點數陣列
type FieldDoubleArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldDoubleArray) TypeExcel() string {
	return "doubleArray"
}

// TypeCpp 取得c++欄位類型
func (this *FieldDoubleArray) TypeCpp() string {
	return "std::vector<double>"
}

// TypeCs 取得c#欄位類型
func (this *FieldDoubleArray) TypeCs() string {
	return "List<double>"
}

// TypeGo 取得go欄位類型
func (this *FieldDoubleArray) TypeGo() string {
	return "[]float64"
}

// Hide 是否隱藏
func (this *FieldDoubleArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldDoubleArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldDoubleArray) Transform(input string) (result interface{}, err error) {
	return util.StringToFloat64Array(input)
}
