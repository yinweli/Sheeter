package core

// FieldEmpty 空欄位
type FieldEmpty struct {
}

// Type 取得excel欄位類型
func (this *FieldEmpty) Type() string {
	return "empty"
}

// IsShow 是否顯示
func (this *FieldEmpty) IsShow() bool {
	return false
}

// IsPkey 是否是主要索引
func (this *FieldEmpty) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *FieldEmpty) ToJsonDefault() interface{} {
	return nil
}

// ToJsonValue 轉換為json值
func (this *FieldEmpty) ToJsonValue(input string) (result interface{}, err error) {
	return nil, nil
}

// ToLuaValue 轉換為lua值
func (this *FieldEmpty) ToLuaValue(input string) (result string, err error) {
	return "", nil
}
