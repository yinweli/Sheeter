package field

import "Sheeter/internal/util"

// LongArray 64位元整數陣列
type LongArray struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *LongArray) TypeExcel() string {
	return "longArray"
}

// TypeCpp 取得c++欄位類型
func (this *LongArray) TypeCpp() string {
	return "std::vector<int64_t>"
}

// TypeCs 取得c#欄位類型
func (this *LongArray) TypeCs() string {
	return "List<long>"
}

// TypeGo 取得go欄位類型
func (this *LongArray) TypeGo() string {
	return "[]int64"
}

// Hide 是否隱藏
func (this *LongArray) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *LongArray) PrimaryKey() bool {
	return false
}

// GetNote 取得註解名稱
func (this *LongArray) GetNote() string {
	return this.Note
}

// GetName 取得欄位名稱
func (this *LongArray) GetName() string {
	return this.Name
}

// GetField 取得欄位類型
func (this *LongArray) GetField() string {
	return this.Field
}

// FillToMetas 寫入到元資料列表
func (this *LongArray) FillToMetas(metas Metas, data string) error {
	values, err := util.StringToInt64Array(data)

	if err != nil {
		return err
	} // if

	metas[this.Name] = values
	return nil
}
