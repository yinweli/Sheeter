package build

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// FieldInt 32位元整數
type FieldInt struct {
}

// Type 取得excel欄位類型
func (this *FieldInt) Type() string {
	return "int"
}

// IsShow 是否顯示
func (this *FieldInt) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldInt) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *FieldInt) ToJsonDefault() interface{} {
	return int64(0)
}

// ToJsonValue 轉換為json值
func (this *FieldInt) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToInt(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}

// ToLuaValue 轉換為lua值
func (this *FieldInt) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToInt(input); err != nil {
		return "", fmt.Errorf("to lua value failed: %w", err)
	} // if

	return input, nil
}
