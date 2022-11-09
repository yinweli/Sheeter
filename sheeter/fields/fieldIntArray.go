package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// IntArray 32位元整數陣列
type IntArray struct {
}

// Type 取得excel欄位類型
func (this *IntArray) Type() string {
	return "intArray"
}

// IsShow 是否顯示
func (this *IntArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *IntArray) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *IntArray) ToTypeCs() string {
	return sheeter.TokenIntCs + sheeter.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *IntArray) ToTypeGo() string {
	return sheeter.TokenArray + sheeter.TokenIntGo
}

// ToTypeProto 取得proto類型字串
func (this *IntArray) ToTypeProto() string {
	return sheeter.TokenRepeated + " " + sheeter.TokenIntProto
}

// ToJsonValue 轉換為json值
func (this *IntArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToIntArray(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
