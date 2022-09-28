package mixeds

import (
	"github.com/yinweli/Sheeter/internal/utils"
)

// String 字串綜合工具
type String struct {
}

// FirstUpper 字串首字母大寫
func (this *String) FirstUpper(input string) string {
	return utils.FirstUpper(input)
}

// FirstLower 字串首字母小寫
func (this *String) FirstLower(input string) string {
	return utils.FirstLower(input)
}
