package util

import "strings"

// FirstUpper 字串首字母大寫
func FirstUpper(input string) string {
	if input == "" {
		return ""
	} // if

	return strings.ToUpper(input[:1]) + input[1:]
}

// FirstLower 字串首字母小寫
func FirstLower(input string) string {
	if input == "" {
		return ""
	} // if

	return strings.ToLower(input[:1]) + input[1:]
}

// VariableCheck 變數名稱檢查
func VariableCheck(input string) bool {
	if len(input) <= 0 { // 變數名稱不能為空
		return false
	} // if

	first := input[0]

	if first >= '0' && first <= '9' { // 變數名稱不能以數字開頭
		return false
	} // if

	for _, itor := range input { // 變數名稱必須是字母, 數字與'_'的組合
		if (itor < 'a' || itor > 'z') && (itor < 'A' || itor > 'Z') && (itor < '0' || itor > '9') && itor != '_' {
			return false
		} // if
	} // for

	return true
}
