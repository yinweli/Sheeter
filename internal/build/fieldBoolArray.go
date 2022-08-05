package build

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// FieldBoolArray 布林值陣列
type FieldBoolArray struct {
}

// Type 取得excel欄位類型
func (this *FieldBoolArray) Type() string {
	return "boolArray"
}

// IsShow 是否顯示
func (this *FieldBoolArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldBoolArray) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *FieldBoolArray) ToJsonDefault() interface{} {
	return []bool{}
}

// ToJsonValue 轉換為json值
func (this *FieldBoolArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToBoolArray(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}

// ToLuaValue 轉換為lua值
func (this *FieldBoolArray) ToLuaValue(input string) (result string, err error) {
	if result, err = util.LuaBoolArray(input); err != nil {
		return "", fmt.Errorf("to lua value failed: %w", err)
	} // if

	return util.LuaWrapperArray(result), nil
}
