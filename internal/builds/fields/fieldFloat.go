package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// Float 32位元浮點數
type Float struct {
}

// Type 取得excel欄位類型
func (this *Float) Type() string {
	return "float"
}

// IsShow 是否顯示
func (this *Float) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Float) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *Float) ToJsonDefault() interface{} {
	return float64(0)
}

// ToJsonValue 轉換為json值
func (this *Float) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToFloat(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
