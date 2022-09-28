package testdata

import (
	"runtime"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal"
)

const UnknownStr = "?????"
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
const SheetName = "Data"
const ProtoNameTest = "test1.proto"
const ProtoDataReal = "real.pbd"

// IsWindows 取得是否在windows下執行
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// GetTestExcel 取得測試excel
func GetTestExcel(path string) *excelize.File {
	excel, err := excelize.OpenFile(path)

	if err != nil {
		return nil
	} // if

	return excel
}

// GetExcelContentReal 取得ExcelNameReal所指的excel檔案轉為物件後的內容
func GetExcelContentReal() interface{} {
	return map[string]interface{}{
		"Datas": map[internal.PkeyType]interface{}{
			1: map[string]interface{}{
				"Name0": 1,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
						{"Name2": 1, "Name3": "a"},
					},
					"Name1": true,
				},
			},
			2: map[string]interface{}{
				"Name0": 2,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
						{"Name2": 2, "Name3": "b"},
					},
					"Name1": false,
				},
			},
			3: map[string]interface{}{
				"Name0": 3,
				"S": map[string]interface{}{
					"A": []map[string]interface{}{
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
						{"Name2": 3, "Name3": "c"},
					},
					"Name1": true,
				},
			},
		},
	}
}

// GetExcelContentEmpty 取得ExcelNameEmpty所指的excel檔案轉為物件後的內容
func GetExcelContentEmpty() interface{} {
	return map[string]interface{}{
		"Datas": map[internal.PkeyType]interface{}{},
	}
}
