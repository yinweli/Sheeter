package core

import "Sheeter/internal/util"

// FieldFloatArray 32位元浮點數陣列
type FieldFloatArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldFloatArray) TypeExcel() string {
	return "floatArray"
}

// TypeCpp 取得c++欄位類型
func (this *FieldFloatArray) TypeCpp() string {
	return "std::vector<float>"
}

// TypeCs 取得c#欄位類型
func (this *FieldFloatArray) TypeCs() string {
	return "List<float>"
}

// TypeGo 取得go欄位類型
func (this *FieldFloatArray) TypeGo() string {
	return "[]float32"
}

// Hide 是否隱藏
func (this *FieldFloatArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldFloatArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldFloatArray) Transform(input string) (result interface{}, err error) {
	return util.StringToFloat32Array(input)
}
