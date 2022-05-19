package util

import "Sheeter/internal/logger"

// Checker 檢查器
type Checker struct {
	result bool // 檢查結果
}

// Add 新增檢查
func (this *Checker) Add(condition bool, error string) {
	if condition == false && error != "" {
		logger.Error(error)
	}

	this.result = this.result && condition

}

// Result 取得檢查結果
func (this *Checker) Result() (result bool) {
	return this.result
}

// NewChecker 建立檢查器
func NewChecker() (checker *Checker) {
	return &Checker{result: true}
}
