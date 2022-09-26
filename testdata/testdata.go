package testdata

import (
	"runtime"

	"github.com/xuri/excelize/v2"
)

const ConfigNameReal = "config.real.yaml"
const ConfigNameFake = "config.fake.yaml"
const ExcelNameReal = "real.xlsx"
const ExcelNameEmpty = "empty.xlsx"
const ExcelNameCleanAll = "excel.clean.all.xlsx"
const ExcelNameCleanField = "excel.clean.field.xlsx"
const ExcelNameInvalidData = "excel.invalid.data.xlsx"
const ExcelNameInvalidField = "excel.invalid.field.xlsx"
const ExcelNameInvalidFile = "excel.invalid.file.xlsx"
const ExcelNameInvalidLayer = "excel.invalid.layer.xlsx"
const ExcelNameInvalidLayout = "excel.invalid.layout.xlsx"
const ExcelNameInvalidPkeyDupl = "excel.invalid.pkey.dupl.xlsx"
const ExcelNameInvalidPkeyZero = "excel.invalid.pkey.zero.xlsx"
const TestProto = "test1.proto"
const SheetName = "Data"
const UnknownStr = "?????"

// GetTestExcel 取得測試excel
func GetTestExcel(path string) *excelize.File {
	excel, err := excelize.OpenFile(path)

	if err != nil {
		return nil
	} // if

	return excel
}

// IsWindows 取得是否在windows下執行
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
