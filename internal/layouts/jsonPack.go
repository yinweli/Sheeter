package layouts

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/internal/utils"
)

// JsonPack 打包json資料
func JsonPack(line *excels.Line, layoutJson *LayoutJson, excludes []string) (json []byte, err error) {
	defer line.Close()
	datas := map[internal.PkeyType]interface{}{}

	for ok := true; ok; ok = line.Next() {
		raws, _ := line.Data()

		if raws == nil {
			break // 碰到空行就結束了
		} // if

		packs, pkey, err := layoutJson.Pack(raws, excludes)

		if err != nil {
			return nil, fmt.Errorf("json pack failed: %w", err)
		} // if

		datas[pkey] = jsonFirstUpper(packs) // 因為轉為proto資料時需要欄位為大寫駝峰, 所以在此轉換
	} // for

	obj := map[string]interface{}{} // 因為轉為proto資料時需要多包一層, 所以json資料也跟著多包一層
	obj[internal.StorerDatas] = datas

	if json, err = utils.JsonMarshal(obj); err != nil {
		return nil, fmt.Errorf("json pack failed: %w", err)
	} // if

	return json, nil
}

// jsonFirstUpper 把json的欄位改為大寫駝峰
func jsonFirstUpper(input map[string]interface{}) (result map[string]interface{}) {
	result = map[string]interface{}{}

	for k, v := range input {
		key := utils.FirstUpper(k)

		switch value := v.(type) {
		case map[string]interface{}:
			result[key] = jsonFirstUpper(value)

		case *[]map[string]interface{}: // 這邊要注意, 由於structor中的陣列元素是指標, 所以這裡也得是指標才行
			array := []map[string]interface{}{}

			for _, itor := range *value {
				array = append(array, jsonFirstUpper(itor))
			} // for

			result[key] = &array // 回存也得是指標

		default:
			result[key] = v
		} // switch
	} // for

	return result
}
