package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// Float 32位元浮點數
type Float struct {
}

// Field 取得excel欄位類型列表
func (this *Float) Field() []string {
	return []string{"float"}
}

// IsPkey 是否是主要索引
func (this *Float) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *Float) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *Float) ToTypeCs() string {
	return sheeter.TypeFloatCs
}

// ToTypeGo 取得go類型字串
func (this *Float) ToTypeGo() string {
	return sheeter.TypeFloatGo
}

// ToJsonValue 轉換為json值
func (this *Float) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat32(input)

	if err != nil {
		return nil, fmt.Errorf("float to json value: %w", err)
	} // if

	return result, nil
}
