package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Long 64位元整數
type Long struct {
}

// Field 取得excel欄位類型列表
func (this *Long) Field() []string {
	return []string{"long"}
}

// IsShow 是否顯示
func (this *Long) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Long) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *Long) ToTypeCs() string {
	return sheeter.TokenLongCs
}

// ToTypeGo 取得go類型字串
func (this *Long) ToTypeGo() string {
	return sheeter.TokenLongGo
}

// ToTypeProto 取得proto類型字串
func (this *Long) ToTypeProto() string {
	return sheeter.TokenLongProto
}

// ToJsonValue 轉換為json值
func (this *Long) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt64(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
