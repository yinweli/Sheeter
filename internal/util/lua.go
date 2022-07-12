package util

import (
	"strings"
)

type LuaObj = map[string]string  // lua物件型態
type LuaObjs = map[string]LuaObj // lua列表型態

// LuaBool 轉換為lua布林值
func LuaBool(input string) (result string, err error) {
	value, err := StrToBool(input)

	if err != nil {
		return "", err
	} // if

	if value {
		return "true", nil
	} else {
		return "false", nil
	} // if
}

// LuaBoolArray 轉換為lua布林值陣列
func LuaBoolArray(input string) (result string, err error) {
	var values []string

	for _, itor := range strings.Split(input, arraySeparator) {
		value, err := LuaBool(itor)

		if err != nil {
			return "", err
		} // if

		values = append(values, value)
	} // for

	return strings.Join(values, arraySeparator), nil
}

// LuaWrapperArray 包裝lua陣列
func LuaWrapperArray(input string) string {
	return "{" + input + "}"
}

// LuaWrapperString 包裝lua字串
func LuaWrapperString(input string) string {
	return "\"" + input + "\""
}
