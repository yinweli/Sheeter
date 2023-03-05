package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// FloatArray 32位元浮點數陣列
type FloatArray struct {
}

// Field 取得excel欄位類型列表
func (this *FloatArray) Field() []string {
	return []string{"floatArray", "[]float", "float[]"}
}

// IsShow 是否顯示
func (this *FloatArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FloatArray) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *FloatArray) ToTypeCs() string {
	return sheeter.TokenFloatCs + sheeter.TokenArray
}

// ToTypeGo 取得go類型字串
func (this *FloatArray) ToTypeGo() string {
	return sheeter.TokenArray + sheeter.TokenFloatGo
}

// ToTypeProto 取得proto類型字串
func (this *FloatArray) ToTypeProto() string {
	return sheeter.TokenRepeated + " " + sheeter.TokenFloatProto
}

// ToJsonValue 轉換為json值
func (this *FloatArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat32Array(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
