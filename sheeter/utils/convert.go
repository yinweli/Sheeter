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

// StrToInt32 字串轉32位元整數
func StrToInt32(input string) (result int32, err error) {
	result64, err := strconv.ParseInt(input, 10, 32)

	if err != nil {
		return 0, fmt.Errorf("str to int failed: %w", err)
	} // if

	return int32(result64), nil
}

// StrToInt32Array 字串轉32位元整數陣列
func StrToInt32Array(input string) (result []int32, err error) {
	for _, itor := range strings.Split(input, sheeter.SeparateArray) {
		value, err := StrToInt32(itor)

		if err != nil {
			return nil, err
		} // if

		result = append(result, value)
	} // for

	return result, nil
}

// StrToInt64 字串轉64位元整數
func StrToInt64(input string) (result int64, err error) {
	result, err = strconv.ParseInt(input, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("str to int failed: %w", err)
	} // if

	return result, nil
}

// StrToInt64Array 字串轉64位元整數陣列
func StrToInt64Array(input string) (result []int64, err error) {
	for _, itor := range strings.Split(input, sheeter.SeparateArray) {
		value, err := StrToInt64(itor)

		if err != nil {
			return nil, err
		} // if

		result = append(result, value)
	} // for

	return result, nil
}

// StrToFloat32 字串轉32位元浮點數
func StrToFloat32(input string) (result float32, err error) {
	result64, err := strconv.ParseFloat(input, 32)

	if err != nil {
		return 0, fmt.Errorf("str to float failed: %w", err)
	} // if

	return float32(result64), nil
}

// StrToFloat32Array 字串轉32位元浮點數陣列
func StrToFloat32Array(input string) (result []float32, err error) {
	for _, itor := range strings.Split(input, sheeter.SeparateArray) {
		value, err := StrToFloat32(itor)

		if err != nil {
			return nil, err
		} // if

		result = append(result, value)
	} // for

	return result, nil
}

// StrToFloat64 字串轉64位元浮點數
func StrToFloat64(input string) (result float64, err error) {
	result64, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, fmt.Errorf("str to float failed: %w", err)
	} // if

	return result64, nil
}

// StrToFloat64Array 字串轉64位元浮點數陣列
func StrToFloat64Array(input string) (result []float64, err error) {
	for _, itor := range strings.Split(input, sheeter.SeparateArray) {
		value, err := StrToFloat64(itor)

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
