package field

import (
	"Sheeter/internal/util"
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
	values, err := util.StringToBoolArray(data)

	if err != nil {
		return err
	} // if

	metas[this.Name] = values
	return nil
}
