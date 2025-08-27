package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Int 32位元整數
type Int struct {
}

// Field 取得excel欄位類型列表
func (this *Int) Field() []string {
	return []string{"int"}
}

// ToTypeCs 取得cs類型字串
func (this *Int) ToTypeCs() string {
	return "Int32"
}

// ToTypeGo 取得go類型字串
func (this *Int) ToTypeGo() string {
	return "int32"
}

// ToJsonValue 轉換為json值
func (this *Int) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt32(input)

	if err != nil {
		return nil, fmt.Errorf("int to json value: %w", err)
	} // if

	return result, nil
}
