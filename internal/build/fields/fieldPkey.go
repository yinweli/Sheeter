package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// Pkey 主要索引
type Pkey struct {
}

// Type 取得excel欄位類型
func (this *Pkey) Type() string {
	return "pkey"
}

// IsShow 是否顯示
func (this *Pkey) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Pkey) IsPkey() bool {
	return true
}

// ToJsonDefault 轉換為json預設值
func (this *Pkey) ToJsonDefault() interface{} {
	return int64(0)
}

// ToJsonValue 轉換為json值
func (this *Pkey) ToJsonValue(input string) (result interface{}, err error) {
	result, err = util.StrToInt(input)

	if err != nil {
		return nil, fmt.Errorf("to json value failed: %w", err)
	} // if

	return result, nil
}
