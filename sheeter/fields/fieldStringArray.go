package fields

import (
	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// StringArray 字串陣列
type StringArray struct {
}

// Field 取得excel欄位類型列表
func (this *StringArray) Field() []string {
	return []string{"stringArray", "[]string", "string[]"}
}

// IsPkey 是否是主要索引
func (this *StringArray) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *StringArray) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *StringArray) ToTypeCs() string {
	return sheeter.TypeStringCs + sheeter.TypeArray
}

// ToTypeGo 取得go類型字串
func (this *StringArray) ToTypeGo() string {
	return sheeter.TypeArray + sheeter.TypeStringGo
}

// ToJsonValue 轉換為json值
func (this *StringArray) ToJsonValue(input string) (result interface{}, err error) {
	return utils.StrToStrArray(input), nil
}
