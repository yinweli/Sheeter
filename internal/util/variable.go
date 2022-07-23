package util

// VariableCheck 變數名稱檢查
func VariableCheck(input string) bool {
	if input == "" { // 變數名稱不能為空
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
