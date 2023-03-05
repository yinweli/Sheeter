package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Bool 布林值
type Bool struct {
}

// Field 取得excel欄位類型列表
func (this *Bool) Field() []string {
	return []string{"bool"}
}

// IsShow 是否顯示
func (this *Bool) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Bool) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *Bool) ToTypeCs() string {
	return sheeter.TokenBoolCs
}

// ToTypeGo 取得go類型字串
func (this *Bool) ToTypeGo() string {
	return sheeter.TokenBoolGo
}

// ToTypeProto 取得proto類型字串
func (this *Bool) ToTypeProto() string {
	return sheeter.TokenBoolProto
}

// ToJsonValue 轉換為json值
func (this *Bool) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToBool(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
