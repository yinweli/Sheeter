package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldTextArray 字串陣列
type FieldTextArray struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldTextArray) TypeExcel() string {
	return "textArray"
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

// ToJsonValue 轉換為json值
func (this *FieldTextArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToStrArray(input), nil
}

// ToLuaValue 轉換為lua值
func (this *FieldTextArray) ToLuaValue(input string) (result string, err error) {
	return util.LuaArrayWrapper(input), nil
}
