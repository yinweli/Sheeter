package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Double 64位元浮點數
type Double struct {
}

// Field 取得excel欄位類型列表
func (this *Double) Field() []string {
	return []string{"double"}
}

// ToTypeCs 取得cs類型字串
func (this *Double) ToTypeCs() string {
	return "double"
}

// ToTypeGo 取得go類型字串
func (this *Double) ToTypeGo() string {
	return "float64"
}

// ToJsonValue 轉換為json值
func (this *Double) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToFloat64(input)

	if err != nil {
		return nil, fmt.Errorf("double to json value: %w", err)
	} // if

	return result, nil
}
