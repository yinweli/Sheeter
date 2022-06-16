package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldBoolArray 布林值陣列
type FieldBoolArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldBoolArray) TypeExcel() string {
	return "boolArray"
}

// TypeCs 取得c#欄位類型
func (this *FieldBoolArray) TypeCs() string {
	return "List<bool>"
}

// TypeGo 取得go欄位類型
func (this *FieldBoolArray) TypeGo() string {
	return "[]bool"
}

// IsShow 是否顯示
func (this *FieldBoolArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldBoolArray) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldBoolArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToBoolArray(input)
}

// ToLuaValue 轉換為lua值
func (this *FieldBoolArray) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToBoolArray(input); err != nil {
		return "", err
	} // if

	return util.LuaArrayWrapper(input), nil
}
