package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
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

// ToTypeCs 取得cs類型字串
func (this *Float) ToTypeCs() string {
	return internal.TokenFloatCs
}

// ToTypeGo 取得go類型字串
func (this *Float) ToTypeGo() string {
	return internal.TokenFloatGo
}

// ToTypeProto 取得proto類型字串
func (this *Float) ToTypeProto() string {
	return internal.TokenFloatProto
}

// ToJsonValue 轉換為json值
func (this *Float) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset {
		return float64(0), nil
	} // if

	result, err = utils.StrToFloat(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
