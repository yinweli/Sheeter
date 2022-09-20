package fields

import (
	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
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

// ToTypeCs 取得cs類型字串
func (this *TextArray) ToTypeCs() string {
	return internal.TokenStringCs + internal.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *TextArray) ToTypeGo() string {
	return internal.TokenArray + internal.TokenStringGo
}

// ToTypeProto 取得proto類型字串
func (this *TextArray) ToTypeProto() string {
	return internal.TokenRepeated + " " + internal.TokenStringProto
}

// ToJsonValue 轉換為json值
func (this *TextArray) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset {
		return []string{}, nil
	} // if

	return utils.StrToStrArray(input), nil
}
