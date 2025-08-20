package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// LongArray 64位元整數陣列
type LongArray struct {
}

// Field 取得excel欄位類型列表
func (this *LongArray) Field() []string {
	return []string{"longArray", "[]long", "long[]"}
}

// ToTypeCs 取得cs類型字串
func (this *LongArray) ToTypeCs() string {
	return "long[]"
}

// ToTypeGo 取得go類型字串
func (this *LongArray) ToTypeGo() string {
	return "[]int64"
}

// ToJsonValue 轉換為json值
func (this *LongArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt64Array(input)

	if err != nil {
		return nil, fmt.Errorf("long array to json value: %w", err)
	} // if

	return result, nil
}
