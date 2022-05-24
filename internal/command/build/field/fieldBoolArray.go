package field

import (
	"Sheeter/internal/util"
)

// BoolArray 布林值陣列
type BoolArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *BoolArray) TypeExcel() string {
	return "boolArray"
}

// TypeCpp 取得c++欄位類型
func (this *BoolArray) TypeCpp() string {
	return "std::vector<bool>"
}

// TypeCs 取得c#欄位類型
func (this *BoolArray) TypeCs() string {
	return "List<bool>"
}

// TypeGo 取得go欄位類型
func (this *BoolArray) TypeGo() string {
	return "[]bool"
}

// Hide 是否隱藏
func (this *BoolArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *BoolArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *BoolArray) Transform(input string) (result interface{}, err error) {
	return util.StringToBoolArray(input)
}
