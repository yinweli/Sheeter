package layouts

import (
	"fmt"
	"strings"

	"github.com/yinweli/Sheeter/v3/sheeter"
	"github.com/yinweli/Sheeter/v3/sheeter/excels"
	"github.com/yinweli/Sheeter/v3/sheeter/utils"
)

// JsonPack 打包json資料, 將會把sheet中的資料, 依據資料布局與排除標籤, 轉換為json格式的位元陣列
func JsonPack(tag string, lineOfData int, sheet *excels.Sheet, layout *Layout) (result []byte, err error) {
	data := map[string]interface{}{}
	fail := strings.Builder{}
	sheet.Nextn(lineOfData)

	for ok := true; ok; ok = sheet.Next() {
		line, _ := sheet.Data()

		if line == nil { // 碰到空行就結束了
			break
		} // if

		if output := utils.At(line, sheeter.IndexOutput); utils.CheckIgnore(output) { // 跳過忽略行
			continue
		} // if

		primary, pack, err := layout.Pack(tag, line)

		if err != nil {
			fail.WriteString(fmt.Sprintf("    line(%v): %v\n", sheet.Line(), err))
			continue
		} // if

		if primary == nil {
			fail.WriteString(fmt.Sprintf("    line(%v): primary nil\n", sheet.Line()))
			continue
		} // if

		if pack == nil {
			continue
		} // if

		key := fmt.Sprintf("%v", primary)

		if _, duplicate := data[key]; duplicate {
			fail.WriteString(fmt.Sprintf("    line(%v): primary duplicate: %v\n", sheet.Line(), key))
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
