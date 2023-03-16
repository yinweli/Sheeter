package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// DoubleArray 64位元浮點數陣列
type DoubleArray struct {
}

// Field 取得excel欄位類型列表
func (this *DoubleArray) Field() []string {
	return []string{"doubleArray", "[]double", "double[]"}
}

// IsPkey 是否是主要索引
func (this *DoubleArray) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *DoubleArray) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *DoubleArray) ToTypeCs() string {
	return sheeter.TypeDoubleCs + sheeter.TypeArray
}

// ToTypeGo 取得go類型字串
func (this *DoubleArray) ToTypeGo() string {
	return sheeter.TypeArray + sheeter.TypeDoubleGo
}

// ToJsonValue 轉換為json值
func (this *DoubleArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat64Array(input)

	if err != nil {
		return nil, fmt.Errorf("double array to json value: %w", err)
	} // if

	return result, nil
}
