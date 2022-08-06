package fields

// Empty 空欄位
type Empty struct {
}

// Type 取得excel欄位類型
func (this *Empty) Type() string {
	return "empty"
}

// IsShow 是否顯示
func (this *Empty) IsShow() bool {
	return false
}

// IsPkey 是否是主要索引
func (this *Empty) IsPkey() bool {
	return false
}

// ToJsonDefault 轉換為json預設值
func (this *Empty) ToJsonDefault() interface{} {
	return nil
}

// ToJsonValue 轉換為json值
func (this *Empty) ToJsonValue(input string) (result interface{}, err error) {
	return nil, nil
}

// ToLuaValue 轉換為lua值
func (this *Empty) ToLuaValue(input string) (result string, err error) {
	return "", nil
}
