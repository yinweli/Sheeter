package core

import (
	"Sheeter/internal/util"
)

// FieldBoolArray 布林值陣列
type FieldBoolArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldBoolArray) TypeExcel() string {
	return "boolArray"
}

// TypeCpp 取得c++欄位類型
func (this *FieldBoolArray) TypeCpp() string {
	return "std::vector<bool>"
}

// TypeCs 取得c#欄位類型
func (this *FieldBoolArray) TypeCs() string {
	return "List<bool>"
}

// TypeGo 取得go欄位類型
func (this *FieldBoolArray) TypeGo() string {
	return "[]bool"
}

// Show 是否顯示
func (this *FieldBoolArray) Show() bool {
	return true
}

// PrimaryKey 是否是主要索引
func (this *FieldBoolArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldBoolArray) Transform(input string) (result interface{}, err error) {
	return util.StringToBoolArray(input)
}
