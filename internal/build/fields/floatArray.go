package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
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

// ToJsonDefault 轉換為json預設值
func (this *FloatArray) ToJsonDefault() interface{} {
	return []float64{}
}

// ToJsonValue 轉換為json值
func (this *FloatArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToFloatArray(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}

// ToLuaValue 轉換為lua值
func (this *FloatArray) ToLuaValue(input string) (result string, err error) {
	if _, err := util.StrToFloatArray(input); err != nil {
		return "", fmt.Errorf("to lua value failed: %w", err)
	} // if

	return util.LuaWrapperArray(input), nil
}
