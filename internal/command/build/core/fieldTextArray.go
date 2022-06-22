package core

import (
	"strings"

	"github.com/yinweli/Sheeter/internal/util"
)

// FieldTextArray 字串陣列
type FieldTextArray struct {
}

// Type 取得excel欄位類型
func (this *FieldTextArray) Type() string {
	return "textArray"
}

// IsShow 是否顯示
func (this *FieldTextArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldTextArray) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *FieldTextArray) ToJsonDefault() interface{} {
	return []string{}
}

// ToJsonValue 轉換為json值
func (this *FieldTextArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToStrArray(input), nil
}

// ToLuaValue 轉換為lua值
func (this *FieldTextArray) ToLuaValue(input string) (result string, err error) {
	return util.LuaArrayWrapper("\"" + strings.Join(util.StrToStrArray(input), "\",\"") + "\""), nil
}
