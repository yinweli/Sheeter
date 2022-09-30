package fields

import (
	"github.com/yinweli/Sheeter/internal"
)

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

// ToTypeCs 取得cs類型字串
func (this *Text) ToTypeCs() string {
	return internal.TokenStringCs
}

// ToTypeGo 取得go類型字串
func (this *Text) ToTypeGo() string {
	return internal.TokenStringGo
}

// ToTypeProto 取得proto類型字串
func (this *Text) ToTypeProto() string {
	return internal.TokenStringProto
}

// ToJsonValue 轉換為json值
func (this *Text) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil
}
