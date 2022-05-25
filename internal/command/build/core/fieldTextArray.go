package core

import "Sheeter/internal/util"

// TextArray 字串陣列
type TextArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *TextArray) TypeExcel() string {
	return "textArray"
}

// TypeCpp 取得c++欄位類型
func (this *TextArray) TypeCpp() string {
	return "std::vector<std::string>"
}

// TypeCs 取得c#欄位類型
func (this *TextArray) TypeCs() string {
	return "List<string>"
}

// TypeGo 取得go欄位類型
func (this *TextArray) TypeGo() string {
	return "[]string"
}

// Hide 是否隱藏
func (this *TextArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *TextArray) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *TextArray) Transform(input string) (result interface{}, err error) {
	return util.StringToStringArray(input), nil
}
