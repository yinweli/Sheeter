package utils

import (
	"strings"
	"unicode"

	"github.com/yinweli/Sheeter/v2/sheeter"
)

// CheckIgnore 忽略檢查
func CheckIgnore(input string) bool {
	return strings.Contains(strings.ToLower(input), sheeter.TokenIgnore)
}

// CheckExcel excel名稱檢查
func CheckExcel(input string) bool {
	if input == "" { // 名稱不能為空
		return false
	} // if

	if unicode.IsDigit(rune(input[0])) { // 名稱不能以數字開頭
		return false
	} // if

	for _, itor := range input { // 名稱必須是字母, 數字與'_'的組合
		if unicode.IsLetter(itor) == false && unicode.IsDigit(itor) == false && itor != '_' {
			return false
		} // if
	} // for

	return true
}

// CheckSheet sheet名稱檢查
func CheckSheet(input string) bool {
	if input == "" { // 名稱不能為空
		return false
	} // if

	for _, itor := range input { // 名稱必須是字母, 數字與'_'的組合
		if unicode.IsLetter(itor) == false && unicode.IsDigit(itor) == false && itor != '_' {
			return false
		} // if
	} // for

	return true
}

// CheckField 欄位名稱檢查
func CheckField(input string) bool {
	return CheckExcel(input)
}

// CheckTag 標籤檢查
func CheckTag(input, tag string) bool {
	if CheckIgnore(tag) {
		return false
	} // if

	for _, itor := range input {
		if strings.ContainsRune(tag, itor) {
			return true
		} // if
	} // for

	return false
}
