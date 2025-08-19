package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// FloatArray 32位元浮點數陣列
type FloatArray struct {
}

// Field 取得excel欄位類型列表
func (this *FloatArray) Field() []string {
	return []string{"floatArray", "[]float", "float[]"}
}

// IsPkey 是否是主要索引
func (this *FloatArray) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *FloatArray) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *FloatArray) ToTypeCs() string {
	return sheeter.TypeFloatCs + sheeter.TypeArray
}

// ToTypeGo 取得go類型字串
func (this *FloatArray) ToTypeGo() string {
	return sheeter.TypeArray + sheeter.TypeFloatGo
}

// ToJsonValue 轉換為json值
func (this *FloatArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat32Array(input)

	if err != nil {
		return nil, fmt.Errorf("float array to json value: %w", err)
	} // if

	return result, nil
}
