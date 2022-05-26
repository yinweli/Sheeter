package testdata

import (
	"bytes"

	"github.com/spf13/cobra"
)

const RealConfig = "real.yaml"
const FakeConfig = "fake.yaml"
const DefectConfig = "defect.yaml"
const UnknownConfig = "????.yaml"
const TestExcel = "test.xlsx"
const TestSheet = "Data"

// BoolString 取得測試字串
func BoolString() string {
	return "true,false,false,true,true,false,true"
}

// BoolArray 取得測試陣列
func BoolArray() []bool {
	return []bool{true, false, false, true, true, false, true}
}

// Int32String 取得測試字串
func Int32String() string {
	return "10,8,6,4,2,1,0,-99"
}

// Int32Array 取得測試陣列
func Int32Array() []int32 {
	return []int32{10, 8, 6, 4, 2, 1, 0, -99}
}

// Int64String 取得測試字串
func Int64String() string {
	return "10,8,6,4,2,1,0,-99"
}

// Int64Array 取得測試陣列
func Int64Array() []int64 {
	return []int64{10, 8, 6, 4, 2, 1, 0, -99}
}

// Float32String 取得測試字串
func Float32String() string {
	return "0.101,0.202,0.303,0.404,0.505,-0.909"
}

// Float32Array 取得測試陣列
func Float32Array() []float32 {
	return []float32{0.101, 0.202, 0.303, 0.404, 0.505, -0.909}
}

// Float64String 取得測試字串
func Float64String() string {
	return "0.000101,0.000202,0.000303,0.000404,0.000505,-0.000909"
}

// Float64Array 取得測試陣列
func Float64Array() []float64 {
	return []float64{0.000101, 0.000202, 0.000303, 0.000404, 0.000505, -0.000909}
}

// StringString 取得測試字串
func StringString() string {
	return "1234,5678,9012,3456,7890,1234"
}

// StringArray 取得測試陣列
func StringArray() []string {
	return []string{"1234", "5678", "9012", "3456", "7890", "1234"}
}

// MockCommand 取得測試命令物件
func MockCommand() (buffer *bytes.Buffer, command *cobra.Command) {
	buffer = &bytes.Buffer{}
	command = &cobra.Command{}
	command.SetOut(buffer)

	return buffer, command
}
