package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// Double 64位元浮點數
type Double struct {
}

// Field 取得excel欄位類型列表
func (this *Double) Field() []string {
	return []string{"double"}
}

// IsPkey 是否是主要索引
func (this *Double) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *Double) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *Double) ToTypeCs() string {
	return sheeter.TypeDoubleCs
}

// ToTypeGo 取得go類型字串
func (this *Double) ToTypeGo() string {
	return sheeter.TypeDoubleGo
}

// ToJsonValue 轉換為json值
func (this *Double) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat64(input)

	if err != nil {
		return nil, fmt.Errorf("double to json value: %w", err)
	} // if

	return result, nil
}
