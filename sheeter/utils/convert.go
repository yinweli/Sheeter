package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yinweli/Sheeter/sheeter"
)

// StrToBool 字串轉布林值
func StrToBool(input string) (result bool, err error) {
	result, err = strconv.ParseBool(input)

	if err != nil {
		return false, fmt.Errorf("str to bool failed: %w", err)
	} // if

	return result, nil
}

// StrToBoolArray 字串轉布林值陣列
func StrToBoolArray(input string) (result []bool, err error) {
	for _, itor := range strings.Split(input, sheeter.SeparateArray) {
		value, err := StrToBool(itor)

		if err != nil {
			return nil, err
		} // if

		result = append(result, value)
	} // for

	return result, nil
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
func StrToIntArray(input string) (result []int64, err error) {
	for _, itor := range strings.Split(input, sheeter.SeparateArray) {
		value, err := StrToInt(itor)

		if err != nil {
			return nil, err
		} // if

		result = append(result, value)
	} // for

	return result, nil
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
func StrToFloatArray(input string) (result []float64, err error) {
	for _, itor := range strings.Split(input, sheeter.SeparateArray) {
		value, err := StrToFloat(itor)

		if err != nil {
			return nil, err
		} // if

		result = append(result, value)
	} // for

	return result, nil
}

// StrToStrArray 字串轉為字串陣列
func StrToStrArray(input string) []string {
	return strings.Split(input, sheeter.SeparateArray)
}
