package testdata

import (
	"runtime"

	"github.com/yinweli/Sheeter/internal"
)

const UnknownStr = "?????"
const ConfigReal = "config_real.yaml"
const ConfigFake = "config_fake.yaml"
const ExcelReal = "real.xlsx"
const ExcelEmpty = "excel_empty.xlsx"
const ExcelSheet = "excel_sheet.xlsx"
const ExcelJsonPack = "excel_json_pack.xlsx"
const ExcelCleanAll = "excel_clean_all.xlsx"
const ExcelCleanField = "excel_clean_field.xlsx"
const ExcelInvalidData = "excel_invalid_data.xlsx"
const ExcelInvalidField = "excel_invalid_field.xlsx"
const ExcelInvalidFile = "excel_invalid_file.xlsx"
const ExcelInvalidLayer = "excel_invalid_layer.xlsx"
const ExcelInvalidLayout = "excel_invalid_layout.xlsx"
const ExcelInvalidPkeyDupl = "excel_invalid_pkey_dupl.xlsx"
const ExcelInvalidPkeyZero = "excel_invalid_pkey_zero.xlsx"
const ExcelInvalidEnum = "excel_invalid_enum.xlsx"
const ExcelInvalidEnumDupl = "excel_invalid_enum_dupl.xlsx"
const ExcelInvalidIndex = "excel_invalid_index.xlsx"
const ExcelInvalidIndexDupl = "excel_invalid_index_dupl.xlsx"
const SheetData = "Data"
const SheetEnum = "Enum"
const ProtoTest = "test1.proto"

// IsWindows 取得是否在windows下執行
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// GetExcelContentEmpty 取得ExcelNameEmpty所指的excel檔案轉為物件後的內容
func GetExcelContentEmpty() interface{} {
	return map[string]interface{}{
		"Datas": map[internal.PkeyType]interface{}{},
	}
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

// GetExcelContentJsonPack 取得ExcelNameJsonPack所指的excel檔案轉為物件後的內容
func GetExcelContentJsonPack(exclude bool) interface{} {
	if exclude {
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
	} else {
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
					"Name4": 1,
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
					"Name4": 2,
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
					"Name4": 3,
				},
			},
		}
	} // if
}
