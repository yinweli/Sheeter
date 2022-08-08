package layouts

// checkName 名稱檢查器
type checkName map[string]bool

// check 名稱檢查
func (this *checkName) check(name string) bool {
	if _, ok := (*this)[name]; ok {
		return false
	} // if

	(*this)[name] = true
	return true
}
