package core

// Text 字串
type Text struct {
}

// TypeExcel 取得excel欄位類型
func (this *Text) TypeExcel() string {
	return "text"
}

// TypeCpp 取得c++欄位類型
func (this *Text) TypeCpp() string {
	return "std::string"
}

// TypeCs 取得c#欄位類型
func (this *Text) TypeCs() string {
	return "string"
}

// TypeGo 取得go欄位類型
func (this *Text) TypeGo() string {
	return "string"
}

// Hide 是否隱藏
func (this *Text) Hide() bool {
	return false
}

// PrimaryKey 是否是主要索引
func (this *Text) PrimaryKey() bool {
	return false
}

// Transform 字串轉換
func (this *Text) Transform(input string) (result interface{}, err error) {
	return input, nil
}
