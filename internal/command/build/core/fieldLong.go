package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldLong 64位元整數
type FieldLong struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldLong) TypeExcel() string {
	return "long"
}

// TypeCs 取得c#欄位類型
func (this *FieldLong) TypeCs() string {
	return "long"
}

// TypeGo 取得go欄位類型
func (this *FieldLong) TypeGo() string {
	return "int64"
}

// IsShow 是否顯示
func (this *FieldLong) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldLong) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldLong) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToInt(input)
}

// ToLuaValue 轉換為lua值
func (this *FieldLong) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToInt(input); err != nil {
		return "", err
	} // if

	return input, nil
}
