package core

import "Sheeter/internal/util"

// IntArray 32位元整數陣列
type IntArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *IntArray) TypeExcel() string {
	return "intArray"
}

// TypeCpp 取得c++欄位類型
func (this *IntArray) TypeCpp() string {
	return "std::vector<int32_t>"
}

// TypeCs 取得c#欄位類型
func (this *IntArray) TypeCs() string {
	return "List<int>"
}

// TypeGo 取得go欄位類型
func (this *IntArray) TypeGo() string {
	return "[]int32"
}

// Hide 是否隱藏
func (this *IntArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *IntArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *IntArray) Transform(input string) (result interface{}, err error) {
	return util.StringToInt32Array(input)
}
