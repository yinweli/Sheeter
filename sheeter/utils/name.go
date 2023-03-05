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

// IsDataSheetName 資料表單名稱檢查
func IsDataSheetName(input string) bool {
	return strings.HasPrefix(input, sheeter.SignData)
}

// IsEnumSheetName 列舉表單名稱檢查
func IsEnumSheetName(input string) bool {
	return strings.HasPrefix(input, sheeter.SignEnum)
}

// RemoveSheetPrefix 移除表單開頭, 不會移除忽略表單開頭, 因為忽略表單應該在處理前期就被忽略掉了
func RemoveSheetPrefix(input string) string {
	if strings.HasPrefix(input, sheeter.SignData) {
		return strings.TrimPrefix(input, sheeter.SignData)
	} // if

	if strings.HasPrefix(input, sheeter.SignEnum) {
		return strings.TrimPrefix(input, sheeter.SignEnum)
	} // if

	return input
}
