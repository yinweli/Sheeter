package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Int 32位元整數
type Int struct {
}

// Type 取得excel欄位類型
func (this *Int) Type() string {
	return "int"
}

// IsShow 是否顯示
func (this *Int) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Int) IsPkey() bool {
	return false
}

// ToTypeCs 取得cs類型字串
func (this *Int) ToTypeCs() string {
	return internal.TokenIntCs
}

// ToTypeGo 取得go類型字串
func (this *Int) ToTypeGo() string {
	return internal.TokenIntGo
}

// ToTypeProto 取得proto類型字串
func (this *Int) ToTypeProto() string {
	return internal.TokenIntProto
}

// ToJsonValue 轉換為json值
func (this *Int) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset {
		return int64(0), nil
	} // if

	result, err = utils.StrToInt(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
