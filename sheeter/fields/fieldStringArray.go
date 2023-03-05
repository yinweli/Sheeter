package fields

import (
	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// StringArray 字串陣列
type StringArray struct {
}

// Field 取得excel欄位類型列表
func (this *StringArray) Field() []string {
	return []string{"stringArray", "[]string", "string[]"}
}

// IsShow 是否顯示
func (this *StringArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *StringArray) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *StringArray) ToTypeCs() string {
	return sheeter.TokenStringCs + sheeter.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *StringArray) ToTypeGo() string {
	return sheeter.TokenArray + sheeter.TokenStringGo
}

// ToTypeProto 取得proto類型字串
func (this *StringArray) ToTypeProto() string {
	return sheeter.TokenRepeated + " " + sheeter.TokenStringProto
}

// ToJsonValue 轉換為json值
func (this *StringArray) ToJsonValue(input string) (result interface{}, err error) {
	return utils.StrToStrArray(input), nil
}
