package build

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// FieldPkey 主要索引
type FieldPkey struct {
}

// Type 取得excel欄位類型
func (this *FieldPkey) Type() string {
	return "pkey"
}

// IsShow 是否顯示
func (this *FieldPkey) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldPkey) IsPkey() bool {
	return true
}

// ToJsonDefault 轉換為json預設值
func (this *FieldPkey) ToJsonDefault() interface{} {
	return int64(0)
}

// ToJsonValue 轉換為json值
func (this *FieldPkey) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToInt(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}

// ToLuaValue 轉換為lua值
func (this *FieldPkey) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToInt(input); err != nil {
		return "", fmt.Errorf("to lua value failed: %w", err)
	} // if

	return input, nil
}
