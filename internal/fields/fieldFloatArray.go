package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/utils"
)

// FloatArray 32位元浮點數陣列
type FloatArray struct {
}

// Type 取得excel欄位類型
func (this *FloatArray) Type() string {
	return "floatArray"
}

// IsShow 是否顯示
func (this *FloatArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FloatArray) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FloatArray) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset {
		return []float64{}, nil
	} // if

	result, err = utils.StrToFloatArray(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
