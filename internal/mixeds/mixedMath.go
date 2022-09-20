package mixeds

import (
	"strconv"
)

// Math 數學綜合工具
type Math struct {
}

// Add 加法
func (this *Math) Add(l, r int) string {
	return strconv.Itoa(l + r)
}

// Sub 減法
func (this *Math) Sub(l, r int) string {
	return strconv.Itoa(l - r)
}

// Mul 乘法
func (this *Math) Mul(l, r int) string {
	return strconv.Itoa(l * r)
}

// Div 除法
func (this *Math) Div(l, r int) string {
	return strconv.Itoa(l / r)
}
