package core

import "Sheeter/internal/util"

// FieldIntArray 32位元整數陣列
type FieldIntArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldIntArray) TypeExcel() string {
	return "intArray"
}

// TypeCpp 取得c++欄位類型
func (this *FieldIntArray) TypeCpp() string {
	return "std::vector<int32_t>"
}

// TypeCs 取得c#欄位類型
func (this *FieldIntArray) TypeCs() string {
	return "List<int>"
}

// TypeGo 取得go欄位類型
func (this *FieldIntArray) TypeGo() string {
	return "[]int32"
}

// Hide 是否隱藏
func (this *FieldIntArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldIntArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldIntArray) Transform(input string) (result interface{}, err error) {
	return util.StringToInt32Array(input)
}
