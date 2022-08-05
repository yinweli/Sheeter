package build

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// FieldFloat 32位元浮點數
type FieldFloat struct {
}

// Type 取得excel欄位類型
func (this *FieldFloat) Type() string {
	return "float"
}

// IsShow 是否顯示
func (this *FieldFloat) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldFloat) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *FieldFloat) ToJsonDefault() interface{} {
	return float64(0)
}

// ToJsonValue 轉換為json值
func (this *FieldFloat) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToFloat(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}

// ToLuaValue 轉換為lua值
func (this *FieldFloat) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToFloat(input); err != nil {
		return "", fmt.Errorf("to lua value failed: %w", err)
	} // if

	return input, nil
}
