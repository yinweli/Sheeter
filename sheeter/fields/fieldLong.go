package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v2/sheeter"
	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// Long 64位元整數
type Long struct {
}

// Field 取得excel欄位類型列表
func (this *Long) Field() []string {
	return []string{"long"}
}

// IsPkey 是否是主要索引
func (this *Long) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *Long) ToPkey() Field {
	return &Lkey{}
}

// ToTypeCs 取得cs類型字串
func (this *Long) ToTypeCs() string {
	return sheeter.TypeLongCs
}

// ToTypeGo 取得go類型字串
func (this *Long) ToTypeGo() string {
	return sheeter.TypeLongGo
}

// ToJsonValue 轉換為json值
func (this *Long) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt64(input)

	if err != nil {
		return nil, fmt.Errorf("long to json value: %w", err)
	} // if

	return result, nil
}
