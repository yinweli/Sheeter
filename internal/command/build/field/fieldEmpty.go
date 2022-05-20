package field

// Empty 空欄位
type Empty struct {
	Data
}

// StyleJson 取得json欄位類型
func (this *Empty) StyleJson() string {
	return "empty"
}

// StyleCpp 取得c++欄位類型
func (this *Empty) StyleCpp() string {
	return ""
}

// StyleCs 取得c#欄位類型
func (this *Empty) StyleCs() string {
	return ""
}

// StyleGo 取得go欄位類型
func (this *Empty) StyleGo() string {
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

// FillToMetas 寫入到元資料列表
func (this *Empty) FillToMetas(metas Metas, value string) error {
	return nil
}
