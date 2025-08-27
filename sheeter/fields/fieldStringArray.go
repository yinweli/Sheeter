package fields

import (
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// StringArray 字串陣列
type StringArray struct {
}

// Field 取得excel欄位類型列表
func (this *StringArray) Field() []string {
	return []string{"stringArray", "[]string", "string[]"}
}

// ToTypeCs 取得cs類型字串
func (this *StringArray) ToTypeCs() string {
	return "String[]"
}

// ToTypeGo 取得go類型字串
func (this *StringArray) ToTypeGo() string {
	return "[]string"
}

// ToJsonValue 轉換為json值
func (this *StringArray) ToJsonValue(input string) (result interface{}, err error) {
	return utils.StrToStrArray(input), nil
}
