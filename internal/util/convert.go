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

func BoolArrayToString(input []bool) string {
	var temps []string

	for _, itor := range input {
		temps = append(temps, strconv.FormatBool(itor))
	} // for

	result := strings.Join(temps, internal.Separator)

	return result
}
