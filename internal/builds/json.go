package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// OutputJson 輸出json
func OutputJson(content *Content) error {
	rows, err := content.GetRows(content.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: output json failed, data line not found", content.StructName())
	} // if

	defer func() { _ = rows.Close() }()
	objs := map[string]interface{}{}

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		packs, pkey, err := content.builder.Pack(datas, false)

		if err != nil {
			return fmt.Errorf("%s: output json failed: %w", content.StructName(), err)
		} // if

		objs[pkey] = packs
	} // for

	if err = util.WriteJson(content.FileJson(), objs, content.Bom); err != nil {
		return fmt.Errorf("%s: output json failed: %w", content.StructName(), err)
	} // if

	return nil
}
