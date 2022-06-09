package testdata

import (
	"github.com/xuri/excelize/v2"
)

const RealConfig = "real.yaml"
const Defect1Config = "config.defect1.yaml"
const Defect2Config = "config.defect2.yaml"
const RealExcel = "real.xlsx"
const Defect1Excel = "excel.defect1.xlsx"
const Defect2Excel = "excel.defect2.xlsx"
const Defect3Excel = "excel.defect3.xlsx"
const Defect4Excel = "excel.defect4.xlsx"
const Defect5Excel = "excel.defect5.xlsx"
const Defect6Excel = "excel.defect6.xlsx"
const Defect7Excel = "excel.defect7.xlsx"
const Defect8Excel = "excel.defect8.xlsx"
const Defect9Excel = "excel.defect9.xlsx"
const Defect10Excel = "excel.defect10.xlsx"
const SheetName = "Data"
const Text = "this a string"

// GetTestExcel 取得測試excel
func GetTestExcel(name string) *excelize.File {
	excel, err := excelize.OpenFile(Path(name))

	if err != nil {
		return nil
	} // if

	return excel
}

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
