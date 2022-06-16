package core

// FieldEmpty 空欄位
type FieldEmpty struct {
}

// TypeExcel 取得excel欄位類型
func (this *FieldEmpty) TypeExcel() string {
	return "empty"
}

// TypeCs 取得c#欄位類型
func (this *FieldEmpty) TypeCs() string {
	return ""
}

// TypeGo 取得go欄位類型
func (this *FieldEmpty) TypeGo() string {
	return ""
}

// IsShow 是否顯示
func (this *FieldEmpty) IsShow() bool {
	return false
}

// IsPkey 是否是主要索引
func (this *FieldEmpty) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *FieldEmpty) ToJsonValue(input string) (result interface{}, err error) {
	return nil, nil
}
