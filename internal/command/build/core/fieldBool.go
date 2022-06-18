package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldBool 布林值
type FieldBool struct {
}

// Type 取得excel欄位類型
func (this *FieldBool) Type() string {
	return "bool"
}

// IsShow 是否顯示
func (this *FieldBool) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldBool) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldBool) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToBool(input)
}

// ToLuaValue 轉換為lua值
func (this *FieldBool) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToBool(input); err != nil {
		return "", err
	} // if

	return input, nil
}
