package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/pkg/util"
)

// writeSchema 輸出json架構
func writeSchema(content *Content) error {
	packs, _, err := content.builder.Pack([]string{}, true)

	if err != nil {
		return fmt.Errorf("%s: write json schema failed: %w", content.ShowName(), err)
	} // if

	if err = util.JsonWrite(content.SchemaPath(), packs, content.Bom); err != nil {
		return fmt.Errorf("%s: write json schema failed: %w", content.ShowName(), err)
	} // if

	return nil
}

// writeJson 輸出json
func writeJson(content *Content) error {
	rows, err := content.getRows(content.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: write json failed, data line not found", content.ShowName())
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
			return fmt.Errorf("%s: write json failed: %w", content.ShowName(), err)
		} // if

		objs[pkey] = packs
	} // for

	if err = util.JsonWrite(content.JsonPath(), objs, content.Bom); err != nil {
		return fmt.Errorf("%s: write json failed: %w", content.ShowName(), err)
	} // if

	return nil
}
