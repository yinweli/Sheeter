package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// Bool 布林值
type Bool struct {
}

// Type 取得excel欄位類型
func (this *Bool) Type() string {
	return "bool"
}

// IsShow 是否顯示
func (this *Bool) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Bool) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *Bool) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset == false {
		result, err = util.StrToBool(input)

		if err != nil {
			return nil, fmt.Errorf("to json value failed: %w", err)
		} // if

		return result, nil
	} else {
		return false, nil
	} // if
}
