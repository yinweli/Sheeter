package core

import "Sheeter/internal/util"

// FieldTextArray 字串陣列
type FieldTextArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldTextArray) TypeExcel() string {
	return "textArray"
}

// TypeCpp 取得c++欄位類型
func (this *FieldTextArray) TypeCpp() string {
	return "std::vector<std::string>"
}

// TypeCs 取得c#欄位類型
func (this *FieldTextArray) TypeCs() string {
	return "List<string>"
}

// TypeGo 取得go欄位類型
func (this *FieldTextArray) TypeGo() string {
	return "[]string"
}

// IsShow 是否顯示
func (this *FieldTextArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldTextArray) IsPkey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldTextArray) Transform(input string) (result interface{}, err error) {
	return util.StrToStrArray(input), nil
}
