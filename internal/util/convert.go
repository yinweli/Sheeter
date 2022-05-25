package util

import (
	"strconv"
	"strings"

	"Sheeter/internal"
)

// StringToBoolArray 字串轉為陣列
func StringToBoolArray(input string) (results []bool, err error) {
	tokens := strings.Split(input, internal.ArraySeparator)

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

	return strings.Join(tokens, internal.ArraySeparator)
}

// StringToInt32Array 字串轉為陣列
func StringToInt32Array(input string) (results []int32, err error) {
	tokens := strings.Split(input, internal.ArraySeparator)

	for _, itor := range tokens {
		value, err := strconv.ParseInt(itor, internal.Decimal, 32)

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
		tokens = append(tokens, strconv.FormatInt(int64(itor), internal.Decimal))
	} // for

	return strings.Join(tokens, internal.ArraySeparator)
}

// StringToInt64Array 字串轉為陣列
func StringToInt64Array(input string) (results []int64, err error) {
	tokens := strings.Split(input, internal.ArraySeparator)

	for _, itor := range tokens {
		value, err := strconv.ParseInt(itor, internal.Decimal, 64)

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
		tokens = append(tokens, strconv.FormatInt(itor, internal.Decimal))
	} // for

	return strings.Join(tokens, internal.ArraySeparator)
}

// StringToFloat32Array 字串轉為陣列
func StringToFloat32Array(input string) (results []float32, err error) {
	tokens := strings.Split(input, internal.ArraySeparator)

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
		tokens = append(tokens, trimFloatString(strconv.FormatFloat(float64(itor), 'f', internal.Precision, 32)))
	} // for

	return strings.Join(tokens, internal.ArraySeparator)
}

// StringToFloat64Array 字串轉為陣列
func StringToFloat64Array(input string) (results []float64, err error) {
	tokens := strings.Split(input, internal.ArraySeparator)

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
		tokens = append(tokens, trimFloatString(strconv.FormatFloat(itor, 'f', internal.Precision, 64)))
	} // for

	return strings.Join(tokens, internal.ArraySeparator)
}

// StringToStringArray 字串轉為陣列
func StringToStringArray(input string) []string {
	return strings.Split(input, internal.ArraySeparator)
}

// StringArrayToString 陣列轉為字串
func StringArrayToString(inputs []string) string {
	return strings.Join(inputs, internal.ArraySeparator)
}

// trimFloatString 去除浮點數字串結尾多餘的0或是'.'
func trimFloatString(input string) string {
	for strings.HasSuffix(input, "0") { // 去除浮點數字串結尾有多餘的0
		input = strings.TrimSuffix(input, "0")
	} // for

	if strings.HasSuffix(input, ".") { // 去除浮點數字串結尾有多餘的'.'
		input = strings.TrimSuffix(input, ".")
	} // if

	return input
}
