package utils

import (
	"strings"
)

// FirstUpper 字串首字母大寫
func FirstUpper(input string) string {
	if input == "" {
		return ""
	} // if

	if len(input) == 1 {
		return strings.ToUpper(input)
	} // if

	return strings.ToUpper(input[:1]) + input[1:]
}

// FirstLower 字串首字母小寫
func FirstLower(input string) string {
	if input == "" {
		return ""
	} // if

	if len(input) == 1 {
		return strings.ToLower(input)
	} // if

	return strings.ToLower(input[:1]) + input[1:]
}

// SnakeToCamel 將蛇形命名法(snake_case)轉換為大寫駝峰命名法(CamelCase)
// 這個函數也處理含有空格的字串，將空格視為單詞分隔符，類似於底線
func SnakeToCamel(input string) string {
	input = strings.ReplaceAll(input, " ", "_") // 把空格改為底線, 讓含有空格的字串也可以獲得相同效果
	result := ""

	for _, itor := range strings.Split(input, "_") {
		if itor != "" {
			result += strings.ToUpper(itor[:1]) + itor[1:]
		} // if
	} // for

	return result
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

// Combine 合併字串列表
func Combine(target []string, input []any) []string {
	for _, itor := range input {
		if path, ok := itor.(string); ok {
			target = append(target, path)
		} // of
	} // if

	return target
}

// GetItem 從列表中取得項目; 如果列表數量比索引值還小, 就傳回空字串
func GetItem(input []string, index int) string {
	if index >= 0 && index < len(input) {
		return input[index]
	} // if

	return ""
}

// GetUnique 從列表中取得不重複項目
func GetUnique(input []string) []string {
	check := make(map[string]bool)
	result := []string{}

	for _, itor := range input {
		if _, ok := check[itor]; ok == false {
			check[itor] = true
			result = append(result, itor)
		} // if
	} // for

	return result
}
