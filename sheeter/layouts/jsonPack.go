package layouts

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// JsonPack 打包json資料, 將會把sheet中的資料, 依據資料布局與排除標籤, 轉換為json格式的位元陣列
func JsonPack(tag string, lineOfData int, sheet *excels.Sheet, layout *Layout) (result []byte, err error) {
	if sheet.Nextn(lineOfData) == false {
		return nil, fmt.Errorf("json pack: data line not found")
	} // if

	data := map[string]interface{}{}
	fail := strings.Builder{}

	for ok := true; ok; ok = sheet.Next() {
		line, _ := sheet.Data()

		if line == nil { // 碰到空行就結束了
			break
		} // if

		pack, pkey, err := layout.Pack(tag, line)

		if err != nil {
			fail.WriteString(fmt.Sprintf("    line(%v): %v\n", sheet.Line(), err))
			continue
		} // if

		if pack == nil {
			continue
		} // if

		if pkey == nil {
			fail.WriteString(fmt.Sprintf("    line(%v): pkey nil\n", sheet.Line()))
			continue
		} // if

		key := fmt.Sprintf("%v", pkey)

		if _, duplicate := data[key]; duplicate {
			fail.WriteString(fmt.Sprintf("    line(%v): pkey duplicate: %v\n", sheet.Line(), key))
			continue
		} // if

		data[key] = pack
	} // for

	if fail.Len() > 0 {
		return nil, fmt.Errorf("json pack:\n%v", fail.String())
	} // if

	if result, err = utils.JsonMarshal(data); err != nil {
		return nil, fmt.Errorf("json pack: %w", err)
	} // if

	return result, nil
}
