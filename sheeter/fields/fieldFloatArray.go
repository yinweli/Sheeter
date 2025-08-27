package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// FloatArray 32位元浮點數陣列
type FloatArray struct {
}

// Field 取得excel欄位類型列表
func (this *FloatArray) Field() []string {
	return []string{"floatArray", "[]float", "float[]"}
}

// ToTypeCs 取得cs類型字串
func (this *FloatArray) ToTypeCs() string {
	return "Single[]"
}

// ToTypeGo 取得go類型字串
func (this *FloatArray) ToTypeGo() string {
	return "[]float32"
}

// ToJsonValue 轉換為json值
func (this *FloatArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat32Array(input)

	if err != nil {
		return nil, fmt.Errorf("float array to json value: %w", err)
	} // if

	return result, nil
}
