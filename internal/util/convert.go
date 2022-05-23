package util

import (
	"strconv"
	"strings"

	"Sheeter/internal"
)

// StringToBoolArray 字串轉為布林值陣列
func StringToBoolArray(input string) (results []bool, err error) {
	tokens := strings.Split(input, internal.Separator)

	for _, itor := range tokens {
		value, err := strconv.ParseBool(itor)

		if err != nil {
			return nil, err
		} // if

		results = append(results, value)
	} // for

	return results, nil
}

// BoolArrayToString 布林值轉為字串
func BoolArrayToString(inputs []bool) string {
	var tokens []string

	for _, itor := range inputs {
		tokens = append(tokens, strconv.FormatBool(itor))
	} // for

	return strings.Join(tokens, internal.Separator)
}
