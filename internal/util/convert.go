package util

import (
	"fmt"
	"strconv"
	"strings"
)

const arraySeparator = "," // 陣列分隔字串

// StrToBool 字串轉布林值
func StrToBool(input string) (result bool, err error) {
	result, err = strconv.ParseBool(input)

	if err != nil {
		return false, fmt.Errorf("str to bool failed: %w", err)
	} // if

	return result, nil
}

// StrToBoolArray 字串轉布林值陣列
func StrToBoolArray(input string) (results []bool, err error) {
	for _, itor := range strings.Split(input, arraySeparator) {
		result, err := StrToBool(itor)

		if err != nil {
			return nil, err
		} // if

		results = append(results, result)
	} // for

	return results, nil
}

// StrToInt 字串轉整數
func StrToInt(input string) (result int64, err error) {
	result, err = strconv.ParseInt(input, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("str to int failed: %w", err)
	} // if

	return result, nil
}

// StrToIntArray 字串轉整數陣列
func StrToIntArray(input string) (results []int64, err error) {
	for _, itor := range strings.Split(input, arraySeparator) {
		result, err := StrToInt(itor)

		if err != nil {
			return nil, err
		} // if

		results = append(results, result)
	} // for

	return results, nil
}

// StrToFloat 字串轉浮點數
func StrToFloat(input string) (result float64, err error) {
	result, err = strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, fmt.Errorf("str to float failed: %w", err)
	} // if

	return result, nil
}

// StrToFloatArray 字串轉浮點數陣列
func StrToFloatArray(input string) (results []float64, err error) {
	for _, itor := range strings.Split(input, arraySeparator) {
		result, err := StrToFloat(itor)

		if err != nil {
			return nil, err
		} // if

		results = append(results, result)
	} // for

	return results, nil
}

// StrToStrArray 字串轉為字串陣列
func StrToStrArray(input string) []string {
	return strings.Split(input, arraySeparator)
}
