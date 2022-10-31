package utils

import (
	"strings"
)

// NewDuplicate 建立重複檢查器
func NewDuplicate() *Duplicate {
	return &Duplicate{
		datas: map[string]bool{},
	}
}

// Duplicate 重複檢查器
type Duplicate struct {
	datas map[string]bool // 資料列表
}

// Check 重複檢查
func (this *Duplicate) Check(value ...string) bool {
	builder := strings.Builder{}

	for _, itor := range value {
		builder.WriteString(itor)
	} // for

	result := builder.String()

	if _, ok := this.datas[result]; ok {
		return false
	} // if

	this.datas[result] = true
	return true
}
