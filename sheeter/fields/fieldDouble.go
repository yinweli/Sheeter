package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Double 64位元浮點數
type Double struct {
}

// Field 取得excel欄位類型列表
func (this *Double) Field() []string {
	return []string{"double"}
}

// IsShow 是否顯示
func (this *Double) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Double) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *Double) ToTypeCs() string {
	return sheeter.TokenDoubleCs
}

// ToTypeGo 取得go類型字串
func (this *Double) ToTypeGo() string {
	return sheeter.TokenDoubleGo
}

// ToTypeProto 取得proto類型字串
func (this *Double) ToTypeProto() string {
	return sheeter.TokenDoubleProto
}

// ToJsonValue 轉換為json值
func (this *Double) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat64(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
