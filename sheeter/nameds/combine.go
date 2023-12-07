package nameds

import (
	"strings"

	"github.com/yinweli/Sheeter/v2/sheeter/utils"
)

// param 組合名稱參數
type param struct {
	excelUpper bool   // excel名稱是否要首字大寫
	excelName  string // excel名稱
	sheetName  string // sheet名稱
	last       string // excel與sheet的結尾字串
	ext        string // 副檔名
}

// combine 取得組合名稱, 規則如下
//   - excel名稱: 首字大小寫可選, 移除底線或是空格, 底線或是空格後首字必為大寫
//   - sheet名稱: 首字必為大寫, 移除底線或是空格, 底線或是空格後首字必為大寫
//   - 若excel名稱與sheet名稱相同, 則只留下excel名稱
func combine(p *param) string {
	excel := utils.FileName(p.excelName)
	excel = utils.SnakeToCamel(excel)
	sheet := utils.SnakeToCamel(p.sheetName)

	if strings.EqualFold(excel, sheet) {
		sheet = ""
	} // if

	if p.excelUpper {
		excel = utils.FirstUpper(excel)
	} else {
		excel = utils.FirstLower(excel)
	} // if

	name := strings.Builder{}
	name.WriteString(excel)
	name.WriteString(sheet)
	name.WriteString(p.last)

	if p.ext != "" {
		name.WriteString(p.ext)
	} // if

	return name.String()
}
