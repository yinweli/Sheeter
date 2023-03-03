package nameds

import (
	"strings"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
)

// params 組合名稱參數
type params struct {
	excelName  string // excel名稱
	excelUpper bool   // excel名稱是否要首字大寫
	sheetName  string // sheet名稱
	sheetUpper bool   // sheet名稱是否要首字大寫
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
}

// combine 取得組合名稱
func combine(params *params) string {
	excel := utils.FileName(params.excelName)

	if params.excelUpper {
		excel = utils.FirstUpper(excel)
	} else {
		excel = utils.FirstLower(excel)
	} // if

	sheet := removeSheetPrefix(params.sheetName)

	if params.sheetUpper {
		sheet = utils.FirstUpper(sheet)
	} else {
		sheet = utils.FirstLower(sheet)
	} // if

	name := excel + sheet + params.last

	if params.ext != "" {
		name += params.ext
	} // if

	return name
}

// removeSheetPrefix 移除表單開頭字元
func removeSheetPrefix(sheet string) string {
	if strings.HasPrefix(sheet, sheeter.SignData) {
		return strings.TrimPrefix(sheet, sheeter.SignData)
	} // if

	if strings.HasPrefix(sheet, sheeter.SignEnum) {
		return strings.TrimPrefix(sheet, sheeter.SignEnum)
	} // if

	return sheet
}
