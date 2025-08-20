package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Long 64位元整數
type Long struct {
}

// Field 取得excel欄位類型列表
func (this *Long) Field() []string {
	return []string{"long"}
}

// ToTypeCs 取得cs類型字串
func (this *Long) ToTypeCs() string {
	return "long"
}

// ToTypeGo 取得go類型字串
func (this *Long) ToTypeGo() string {
	return "int64"
}

// ToJsonValue 轉換為json值
func (this *Long) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt64(input)

	if err != nil {
		return nil, fmt.Errorf("long to json value: %w", err)
	} // if

	return result, nil
}
