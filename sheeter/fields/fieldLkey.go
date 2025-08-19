package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// Lkey 主要64位元整數索引
type Lkey struct {
}

// Field 取得excel欄位類型列表
func (this *Lkey) Field() []string {
	return []string{"lkey"}
}

// IsPkey 是否是主要索引
func (this *Lkey) IsPkey() bool {
	return true
}

// ToPkey 取得主要索引類型
func (this *Lkey) ToPkey() Field {
	return this
}

// ToTypeCs 取得cs類型字串
func (this *Lkey) ToTypeCs() string {
	return sheeter.TypeLkeyCs
}

// ToTypeGo 取得go類型字串
func (this *Lkey) ToTypeGo() string {
	return sheeter.TypeLkeyGo
}

// ToJsonValue 轉換為json值
func (this *Lkey) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt64(input)

	if err != nil {
		return nil, fmt.Errorf("lkey to json value: %w", err)
	} // if

	return result, nil
}
