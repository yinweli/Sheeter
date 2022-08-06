package fields

import "github.com/yinweli/Sheeter/internal/util"

// Text 字串
type Text struct {
}

// Type 取得excel欄位類型
func (this *Text) Type() string {
	return "text"
}

// IsShow 是否顯示
func (this *Text) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Text) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *Text) ToJsonDefault() interface{} {
	return ""
}

// ToJsonValue 轉換為json值
func (this *Text) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil
}

// ToLuaValue 轉換為lua值
func (this *Text) ToLuaValue(input string) (result string, err error) {
	return util.LuaWrapperString(input), nil
}
