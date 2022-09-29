package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// BoolArray 布林值陣列
type BoolArray struct {
}

// Type 取得excel欄位類型
func (this *BoolArray) Type() string {
	return "boolArray"
}

// IsShow 是否顯示
func (this *BoolArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *BoolArray) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *BoolArray) ToTypeCs() string {
	return internal.TokenBoolCs + internal.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *BoolArray) ToTypeGo() string {
	return internal.TokenArray + internal.TokenBoolGo
}

// ToTypeProto 取得proto類型字串
func (this *BoolArray) ToTypeProto() string {
	return internal.TokenRepeated + " " + internal.TokenBoolProto
}

// ToJsonValue 轉換為json值
func (this *BoolArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToBoolArray(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
