package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// IntArray 32位元整數陣列
type IntArray struct {
}

// Type 取得excel欄位類型
func (this *IntArray) Type() string {
	return "intArray"
}

// IsShow 是否顯示
func (this *IntArray) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *IntArray) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *IntArray) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset {
		return []int64{}, nil
	} // if

	result, err = util.StrToIntArray(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}