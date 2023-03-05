package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// Pkey 主要索引
type Pkey struct {
}

// Field 取得excel欄位類型列表
func (this *Pkey) Field() []string {
	return []string{"pkey"}
}

// IsShow 是否顯示
func (this *Pkey) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Pkey) IsPkey() bool {
	return true
}

// ToTypeCs 取得cs類型字串
func (this *Pkey) ToTypeCs() string {
	return sheeter.TokenPkeyCs
}

// ToTypeGo 取得go類型字串
func (this *Pkey) ToTypeGo() string {
	return sheeter.TokenPkeyGo
}

// ToTypeProto 取得proto類型字串
func (this *Pkey) ToTypeProto() string {
	return sheeter.TokenPkeyProto
}

// ToJsonValue 轉換為json值
func (this *Pkey) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
