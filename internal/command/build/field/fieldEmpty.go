package field

// Empty 空欄位
type Empty struct {
	Data
}

// TypeExcel 取得excel欄位類型
func (this *Empty) TypeExcel() string {
	return "empty"
}

// TypeCpp 取得c++欄位類型
func (this *Empty) TypeCpp() string {
	return ""
}

// TypeCs 取得c#欄位類型
func (this *Empty) TypeCs() string {
	return ""
}

// TypeGo 取得go欄位類型
func (this *Empty) TypeGo() string {
	return ""
}

// Hide 是否隱藏
func (this *Empty) Hide() bool {
	return true
}

// PrimaryKey 是否是主要索引
func (this *Empty) PrimaryKey() bool {
	return false
}

// GetNote 取得註解名稱
func (this *Empty) GetNote() string {
	return this.Note
}

// GetName 取得欄位名稱
func (this *Empty) GetName() string {
	return this.Name
}

// GetField 取得欄位類型
func (this *Empty) GetField() string {
	return this.Field
}

// FillToMetas 寫入到元資料列表
func (this *Empty) FillToMetas(metas Metas, data string) error {
	return nil
}
