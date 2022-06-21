package core

import (
	"github.com/yinweli/Sheeter/internal/util"
)

// FieldFloatArray 32位元浮點數陣列
type FieldFloatArray struct {
}

// Type 取得excel欄位類型
func (this *FieldFloatArray) Type() string {
	return "floatArray"
}

// IsShow 是否顯示
func (this *FieldFloatArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldFloatArray) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldFloatArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToFloatArray(input)
}

// ToLuaValue 轉換為lua值
func (this *FieldFloatArray) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToFloatArray(input); err != nil {
		return "", err
	} // if

	return util.LuaArrayWrapper(input), nil
}
