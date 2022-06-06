package core

// FieldText 字串
type FieldText struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldText) TypeExcel() string {
	return "text"
}

// TypeCpp 取得c++欄位類型
func (this *FieldText) TypeCpp() string {
	return "std::string"
}

// TypeCs 取得c#欄位類型
func (this *FieldText) TypeCs() string {
	return "string"
}

// TypeGo 取得go欄位類型
func (this *FieldText) TypeGo() string {
	return "string"
}

// IsShow 是否顯示
func (this *FieldText) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *FieldText) IsPkey() bool {
	return false
}

// Transform 字串轉換
func (this *FieldText) Transform(input string) (result interface{}, err error) {
	return input, nil
}
