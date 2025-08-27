package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Float 32位元浮點數
type Float struct {
}

// Field 取得excel欄位類型列表
func (this *Float) Field() []string {
	return []string{"float"}
}

// ToTypeCs 取得cs類型字串
func (this *Float) ToTypeCs() string {
	return "Single"
}

// ToTypeGo 取得go類型字串
func (this *Float) ToTypeGo() string {
	return "float32"
}

// ToJsonValue 轉換為json值
func (this *Float) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat32(input)

	if err != nil {
		return nil, fmt.Errorf("float to json value: %w", err)
	} // if

	return result, nil
}
