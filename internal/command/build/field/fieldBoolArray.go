package field

import (
	"strconv"
	"strings"

	"Sheeter/internal"
)

// BoolArray 布林值陣列
type BoolArray struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *BoolArray) TypeExcel() string {
	return "boolArray"
}

// TypeCpp 取得c++欄位類型
func (this *BoolArray) TypeCpp() string {
	return "std::vector<bool>"
}

// TypeCs 取得c#欄位類型
func (this *BoolArray) TypeCs() string {
	return "List<bool>"
}

// TypeGo 取得go欄位類型
func (this *BoolArray) TypeGo() string {
	return "[]bool"
}

// Hide 是否隱藏
func (this *BoolArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *BoolArray) PrimaryKey() bool {
	return false
}

// FillToMetas 寫入到元資料列表
func (this *BoolArray) FillToMetas(metas Metas, data string) error {
	var values []bool

	for _, itor := range strings.Split(data, internal.Separator) {
		value, err := strconv.ParseBool(itor)

		if err != nil {
			return err
		} // if

		values = append(values, value)
	} // for

	metas[this.Name] = values
	return nil
}
