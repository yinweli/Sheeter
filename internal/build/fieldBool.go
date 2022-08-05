package build

import (
	"fmt"

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

// ToJsonDefault 轉換為json預設值
func (this *FieldBool) ToJsonDefault() interface{} {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldBool) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToBool(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}

// ToLuaValue 轉換為lua值
func (this *FieldBool) ToLuaValue(input string) (result string, err error) {
	if result, err = util.LuaBool(input); err != nil {
		return "", fmt.Errorf("to lua value failed: %w", err)
	} // if

	return result, nil
}
