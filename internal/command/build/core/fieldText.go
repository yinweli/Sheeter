package core

// FieldText 字串
type FieldText struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldText) TypeExcel() string {
	return "text"
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

// ToJsonValue 轉換為json值
func (this *FieldText) ToJsonValue(input string) (result interface{}, err error) {
	return input, nil
}

// ToLuaValue 轉換為lua值
func (this *FieldText) ToLuaValue(input string) (result string, err error) {
	return input, nil
}
