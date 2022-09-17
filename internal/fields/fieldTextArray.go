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
	return internal.TokenString + internal.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *TextArray) ToTypeGo() string {
	return internal.TokenArray + internal.TokenString
}

// ToJsonValue 轉換為json值
func (this *TextArray) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset {
		return []string{}, nil
	} // if

	return utils.StrToStrArray(input), nil
}
