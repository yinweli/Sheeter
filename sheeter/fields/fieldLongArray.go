package fields

import (
	"fmt"

	"github.com/yinweli/Sheeter/v2/sheeter/utils"

	"github.com/yinweli/Sheeter/v2/sheeter"
)

// LongArray 64位元整數陣列
type LongArray struct {
}

// Field 取得excel欄位類型列表
func (this *LongArray) Field() []string {
	return []string{"longArray", "[]long", "long[]"}
}

// IsPkey 是否是主要索引
func (this *LongArray) IsPkey() bool {
	return false
}

// ToPkey 取得主要索引類型
func (this *LongArray) ToPkey() Field {
	return nil
}

// ToTypeCs 取得cs類型字串
func (this *LongArray) ToTypeCs() string {
	return sheeter.TypeLongCs + sheeter.TypeArray
}

// ToTypeGo 取得go類型字串
func (this *LongArray) ToTypeGo() string {
	return sheeter.TypeArray + sheeter.TypeLongGo
}

// ToJsonValue 轉換為json值
func (this *LongArray) ToJsonValue(input string) (result interface{}, err error) {
	result, err = utils.StrToInt64Array(input)

	if err != nil {
		return nil, fmt.Errorf("long array to json value: %w", err)
	} // if

	return result, nil
}
