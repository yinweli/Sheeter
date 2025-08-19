package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// IntArray 32位元整數陣列
type IntArray struct {
}

// Field 取得excel欄位類型列表
func (this *IntArray) Field() []string {
	return []string{"intArray", "[]int", "int[]"}
}

// IsPkey 是否是主要索引
func (this *IntArray) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *IntArray) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *IntArray) ToTypeCs() string {
	return sheeter.TypeIntCs + sheeter.TypeArray
}

// ToTypeGo 取得go類型字串
func (this *IntArray) ToTypeGo() string {
	return sheeter.TypeArray + sheeter.TypeIntGo
}

// ToJsonValue 轉換為json值
func (this *IntArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt32Array(input)

	if err != nil {
		return nil, fmt.Errorf("int array to json value: %w", err)
	} // if

	return result, nil
}
