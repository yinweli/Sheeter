package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Bool 布林值
type Bool struct {
}

// Field 取得excel欄位類型列表
func (this *Bool) Field() []string {
	return []string{"bool"}
}

// IsPkey 是否是主要索引
func (this *Bool) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *Bool) ToPkey() Field {
	return &Pkey{}
}

// ToTypeCs 取得cs類型字串
func (this *Bool) ToTypeCs() string {
	return sheeter.TypeBoolCs
}

// ToTypeGo 取得go類型字串
func (this *Bool) ToTypeGo() string {
	return sheeter.TypeBoolGo
}

// ToJsonValue 轉換為json值
func (this *Bool) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToBool(input)

	if err != nil {
		return nil, fmt.Errorf("bool to json value: %w", err)
	} // if

	return result, nil
}
