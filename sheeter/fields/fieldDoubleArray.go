package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// DoubleArray 64位元浮點數陣列
type DoubleArray struct {
}

// Field 取得excel欄位類型列表
func (this *DoubleArray) Field() []string {
	return []string{"doubleArray", "[]double", "double[]"}
}

// IsShow 是否顯示
func (this *DoubleArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *DoubleArray) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *DoubleArray) ToTypeCs() string {
	return sheeter.TokenDoubleCs + sheeter.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *DoubleArray) ToTypeGo() string {
	return sheeter.TokenArray + sheeter.TokenDoubleGo
}

// ToTypeProto 取得proto類型字串
func (this *DoubleArray) ToTypeProto() string {
	return sheeter.TokenRepeated + " " + sheeter.TokenDoubleProto
}

// ToJsonValue 轉換為json值
func (this *DoubleArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat64Array(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
