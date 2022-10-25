package fields

import (
	"github.com/yinweli/Sheeter/internal"
)

// String 字串
type String struct {
}

// Type 取得excel欄位類型
func (this *String) Type() string {
	return "string"
}

// IsShow 是否顯示
func (this *String) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *String) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *String) ToTypeCs() string {
	return internal.TokenStringCs
}

// ToTypeGo 取得go類型字串
func (this *String) ToTypeGo() string {
	return internal.TokenStringGo
}

// ToTypeProto 取得proto類型字串
func (this *String) ToTypeProto() string {
	return internal.TokenStringProto
}

// ToJsonValue 轉換為json值
func (this *String) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil
}
