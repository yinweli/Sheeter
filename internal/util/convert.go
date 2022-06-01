package util

import (
	"strconv"
	"strings"
)

const arraySeparator = "," // 陣列分隔字串
const precision = 6        // 小數點精度

// StringToBoolArray 字串轉為陣列
func StringToBoolArray(input string) (results []bool, err error) {
	tokens := strings.Split(input, arraySeparator)

	for _, itor := range tokens {
		value, err := strconv.ParseBool(itor)

		if err != nil {
			return nil, err
		} // if

		results = append(results, value)
	} // for

	return results, nil
}

// BoolArrayToString 陣列轉為字串
func BoolArrayToString(inputs []bool) string {
	var tokens []string

	for _, itor := range inputs {
		tokens = append(tokens, strconv.FormatBool(itor))
	} // for

	return strings.Join(tokens, arraySeparator)
}

// StringToInt32Array 字串轉為陣列
func StringToInt32Array(input string) (results []int32, err error) {
	tokens := strings.Split(input, arraySeparator)

	for _, itor := range tokens {
		value, err := strconv.ParseInt(itor, 10, 32)

		if err != nil {
			return nil, err
		} // if

		results = append(results, int32(value))
	} // for

	return results, nil
}

// Int32ArrayToString 陣列轉為字串
func Int32ArrayToString(inputs []int32) string {
	var tokens []string

	for _, itor := range inputs {
		tokens = append(tokens, strconv.FormatInt(int64(itor), 10))
	} // for

	return strings.Join(tokens, arraySeparator)
}

// StringToInt64Array 字串轉為陣列
func StringToInt64Array(input string) (results []int64, err error) {
	tokens := strings.Split(input, arraySeparator)

	for _, itor := range tokens {
		value, err := strconv.ParseInt(itor, 10, 64)

		if err != nil {
			return nil, err
		} // if

		results = append(results, value)
	} // for

	return results, nil
}

// Int64ArrayToString 陣列轉為字串
func Int64ArrayToString(inputs []int64) string {
	var tokens []string

	for _, itor := range inputs {
		tokens = append(tokens, strconv.FormatInt(itor, 10))
	} // for

	return strings.Join(tokens, arraySeparator)
}

// StringToFloat32Array 字串轉為陣列
func StringToFloat32Array(input string) (results []float32, err error) {
	tokens := strings.Split(input, arraySeparator)

	for _, itor := range tokens {
		value, err := strconv.ParseFloat(itor, 32)

		if err != nil {
			return nil, err
		} // if

		results = append(results, float32(value))
	} // for

	return results, nil
}

// Float32ArrayToString 陣列轉為字串
func Float32ArrayToString(inputs []float32) string {
	var tokens []string

	for _, itor := range inputs {
		tokens = append(tokens, TrimFloatString(strconv.FormatFloat(float64(itor), 'f', precision, 32)))
	} // for

	return strings.Join(tokens, arraySeparator)
}

// StringToFloat64Array 字串轉為陣列
func StringToFloat64Array(input string) (results []float64, err error) {
	tokens := strings.Split(input, arraySeparator)

	for _, itor := range tokens {
		value, err := strconv.ParseFloat(itor, 64)

		if err != nil {
			return nil, err
		} // if

		results = append(results, value)
	} // for

	return results, nil
}

// Float64ArrayToString 陣列轉為字串
func Float64ArrayToString(inputs []float64) string {
	var tokens []string

	for _, itor := range inputs {
		tokens = append(tokens, TrimFloatString(strconv.FormatFloat(itor, 'f', precision, 64)))
	} // for

	return strings.Join(tokens, arraySeparator)
}

// StringToStringArray 字串轉為陣列
func StringToStringArray(input string) []string {
	return strings.Split(input, arraySeparator)
}

// StringArrayToString 陣列轉為字串
func StringArrayToString(inputs []string) string {
	return strings.Join(inputs, arraySeparator)
}

// TrimFloatString 去除浮點數字串結尾多餘的0或是'.'
func TrimFloatString(input string) string {
	for strings.HasSuffix(input, "0") { // 去除浮點數字串結尾有多餘的0
		input = strings.TrimSuffix(input, "0")
	} // for

	if strings.HasSuffix(input, ".") { // 去除浮點數字串結尾有多餘的'.'
		input = strings.TrimSuffix(input, ".")
	} // if

	return input
}

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
