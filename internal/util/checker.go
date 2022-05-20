package util

// Checker 檢查器
type Checker struct {
	errors []string // 錯誤訊息列表
}

// Add 新增檢查
func (this *Checker) Add(condition bool, error string) {
	if condition == false {
		this.errors = append(this.errors, error)
	} // if
}

// Result 取得檢查結果
func (this *Checker) Result() (result bool) {
	return len(this.errors) <= 0
}

// Errors 取得錯誤訊息列表
func (this *Checker) Errors() (errors []string) {
	return this.errors
}

// NewChecker 建立檢查器
func NewChecker() (checker *Checker) {
	return &Checker{errors: []string{}}
}
