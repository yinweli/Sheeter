package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// BoolArray 布林值陣列
type BoolArray struct {
}

// Field 取得excel欄位類型列表
func (this *BoolArray) Field() []string {
	return []string{"boolArray", "[]bool", "bool[]"}
}

// IsPkey 是否是主要索引
func (this *BoolArray) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *BoolArray) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *BoolArray) ToTypeCs() string {
	return sheeter.TypeBoolCs + sheeter.TypeArray
}

// ToTypeGo 取得go類型字串
func (this *BoolArray) ToTypeGo() string {
	return sheeter.TypeArray + sheeter.TypeBoolGo
}

// ToJsonValue 轉換為json值
func (this *BoolArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToBoolArray(input)

	if err != nil {
		return nil, fmt.Errorf("bool array to json value: %w", err)
	} // if

	return result, nil
}
