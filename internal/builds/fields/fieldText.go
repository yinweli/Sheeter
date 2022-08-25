package fields

// Text 字串
type Text struct {
}

// Type 取得excel欄位類型
func (this *Text) Type() string {
	return "text"
}

// IsShow 是否顯示
func (this *Text) IsShow() bool {
	return true
}

// IsPkey 是否是主要索引
func (this *Text) IsPkey() bool {
	return false
}

// ToJsonValue 轉換為json值
func (this *Text) ToJsonValue(input string, preset bool) (result interface{}, err error) {
	if preset {
		return "", nil
	} // if

	return input, nil
}
