package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// LongArray 64位元整數陣列
type LongArray struct {
}

// Field 取得excel欄位類型列表
func (this *LongArray) Field() []string {
	return []string{"longArray", "[]long", "long[]"}
}

// IsShow 是否顯示
func (this *LongArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *LongArray) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *LongArray) ToTypeCs() string {
	return sheeter.TokenLongCs + sheeter.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *LongArray) ToTypeGo() string {
	return sheeter.TokenArray + sheeter.TokenLongGo
}

// ToTypeProto 取得proto類型字串
func (this *LongArray) ToTypeProto() string {
	return sheeter.TokenRepeated + " " + sheeter.TokenLongProto
}

// ToJsonValue 轉換為json值
func (this *LongArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt64Array(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
