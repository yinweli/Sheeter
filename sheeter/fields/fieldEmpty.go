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

// ToTypeCs 取得cs類型字串
func (this *Empty) ToTypeCs() string {
	return ""
}

// ToTypeGo 取得go類型字串
func (this *Empty) ToTypeGo() string {
	return ""
}

// ToTypeProto 取得proto類型字串
func (this *Empty) ToTypeProto() string {
	return ""
}

// ToJsonValue 轉換為json值
func (this *Empty) ToJsonValue(input string) (result interface{}, err error) {
	return nil, nil
}
