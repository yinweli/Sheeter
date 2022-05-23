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

// FillToJsons 寫入到json列表
func (this *Empty) FillToJsons(jsons Jsons, value string) error {
	return nil
}
