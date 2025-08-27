package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// DoubleArray 64位元浮點數陣列
type DoubleArray struct {
}

// Field 取得excel欄位類型列表
func (this *DoubleArray) Field() []string {
	return []string{"doubleArray", "[]double", "double[]"}
}

// ToTypeCs 取得cs類型字串
func (this *DoubleArray) ToTypeCs() string {
	return "Double[]"
}

// ToTypeGo 取得go類型字串
func (this *DoubleArray) ToTypeGo() string {
	return "[]float64"
}

// ToJsonValue 轉換為json值
func (this *DoubleArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat64Array(input)

	if err != nil {
		return nil, fmt.Errorf("double array to json value: %w", err)
	} // if

	return result, nil
}
