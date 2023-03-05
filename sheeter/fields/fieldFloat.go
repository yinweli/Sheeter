package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Float 32位元浮點數
type Float struct {
}

// Field 取得excel欄位類型列表
func (this *Float) Field() []string {
	return []string{"float"}
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
	return sheeter.TokenFloatCs
}

// ToTypeGo 取得go類型字串
func (this *Float) ToTypeGo() string {
	return sheeter.TokenFloatGo
}

// ToTypeProto 取得proto類型字串
func (this *Float) ToTypeProto() string {
	return sheeter.TokenFloatProto
}

// ToJsonValue 轉換為json值
func (this *Float) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
