package utils

import (
	"strings"
)

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

// AllSame 檢查字串是否由相同字元組成
func AllSame(input string) bool {
	for _, itor := range input {
		if itor != int32(input[0]) {
			return false
		} // if
	} // for

	return true
}

// GetItem 從列表中取得項目
func GetItem(lists []string, i int) string {
	if i >= 0 && i < len(lists) { // 列表的數量可能因為空白格的關係會短缺, 所以要檢查一下
		return lists[i]
	} // if

	return ""
}
