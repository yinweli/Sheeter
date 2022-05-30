package core

// FieldEmpty 空欄位
type FieldEmpty struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldEmpty) TypeExcel() string {
	return "empty"
}

// TypeCpp 取得c++欄位類型
func (this *FieldEmpty) TypeCpp() string {
	return ""
}

// TypeCs 取得c#欄位類型
func (this *FieldEmpty) TypeCs() string {
	return ""
}

// TypeGo 取得go欄位類型
func (this *FieldEmpty) TypeGo() string {
	return ""
}

// Show 是否顯示
func (this *FieldEmpty) Show() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *FieldEmpty) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldEmpty) Transform(input string) (result interface{}, err error) {
	return nil, nil
}
