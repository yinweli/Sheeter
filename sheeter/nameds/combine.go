package nameds

import (
	"strings"

	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// params 組合名稱參數
type params struct {
	excelUpper bool   // excel名稱是否要首字大寫
	excelName  string // excel名稱
	sheetName  string // sheet名稱
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
}

// combine 取得組合名稱, 規則如下
//   - excel名稱: 首字大小寫可選, 移除底線, 底線後首字必為大寫
//   - sheet名稱: 首字必為大寫, 移除底線, 底線後首字必為大寫
//   - 若excel名稱與sheet名稱相同, 則只留下excel名稱
func combine(params *params) string {
	excel := utils.FileName(params.excelName)
	excel = utils.SnakeToCamel(excel)
	sheet := utils.SnakeToCamel(params.sheetName)

	if strings.EqualFold(excel, sheet) {
		sheet = ""
	} // if

	if params.excelUpper {
		excel = utils.FirstUpper(excel)
	} else {
		excel = utils.FirstLower(excel)
	} // if

	name := excel + sheet + params.last

	if params.ext != "" {
		name += params.ext
	} // if

	return name
}
