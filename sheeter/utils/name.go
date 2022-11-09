package utils

import (
	"strings"

	"github.com/yinweli/Sheeter/sheeter"
)

// NameCheck 名稱檢查
func NameCheck(input string) bool {
	if input == "" { // 名稱不能為空
		return false
	} // if

	if input[0] >= '0' && input[0] <= '9' { // 名稱不能以數字開頭
		return false
	} // if

	for _, itor := range input { // 名稱必須是字母, 數字與'_'的組合
		if (itor < 'a' || itor > 'z') && (itor < 'A' || itor > 'Z') && (itor < '0' || itor > '9') && itor != '_' {
			return false
		} // if
	} // for

	return true
}

// NameKeywords 關鍵字檢查
func NameKeywords(input string) bool {
	for _, itor := range sheeter.Keywords {
		if strings.EqualFold(input, itor) {
			return false
		} // if
	} // for

	return true
}
