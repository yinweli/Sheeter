package field

import "Sheeter/internal/util"

// IntArray 32位元整數陣列
type IntArray struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *IntArray) TypeExcel() string {
	return "intArray"
}

// TypeCpp 取得c++欄位類型
func (this *IntArray) TypeCpp() string {
	return "std::vector<int32_t>"
}

// TypeCs 取得c#欄位類型
func (this *IntArray) TypeCs() string {
	return "List<int>"
}

// TypeGo 取得go欄位類型
func (this *IntArray) TypeGo() string {
	return "[]int32"
}

// Hide 是否隱藏
func (this *IntArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *IntArray) PrimaryKey() bool {
	return false
}

// FillToMetas 寫入到元資料列表
func (this *IntArray) FillToMetas(metas Metas, data string) error {
	values, err := util.StringToInt32Array(data)

	if err != nil {
		return err
	} // if

	metas[this.Name] = values
	return nil
}
