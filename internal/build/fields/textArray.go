package fields

import (
	"strings"

	"github.com/yinweli/Sheeter/internal/util"
)

// TextArray 字串陣列
type TextArray struct {
}

// Type 取得excel欄位類型
func (this *TextArray) Type() string {
	return "textArray"
}

// IsShow 是否顯示
func (this *TextArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *TextArray) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *TextArray) ToJsonDefault() interface{} {
	return []string{}
}

// ToJsonValue 轉換為json值
func (this *TextArray) ToJsonValue(input string) (result interface{}, err error) {
	return util.StrToStrArray(input), nil
}

// ToLuaValue 轉換為lua值
func (this *TextArray) ToLuaValue(input string) (result string, err error) {
	return util.LuaWrapperArray(util.LuaWrapperString(strings.Join(util.StrToStrArray(input), "\",\""))), nil
}
