package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// writeSchema 輸出json架構
func writeSchema(content *Content) error {
	objs := map[string]interface{}{}
	packs, pkey, err := content.builder.Pack([]string{}, true)

	if err != nil {
		return fmt.Errorf("%s: write json schema failed: %w", content.TargetName(), err)
	} // if

	objs[pkey] = packs

	if err = util.JsonWrite(content.SchemaFilePath(), objs, content.Bom); err != nil {
		return fmt.Errorf("%s: write json schema failed: %w", content.TargetName(), err)
	} // if

	return nil
}

// writeJson 輸出json
func writeJson(content *Content) error {
	rows, err := content.getRows(content.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: write json failed, data line not found", content.TargetName())
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
			return fmt.Errorf("%s: write json failed: %w", content.TargetName(), err)
		} // if

		objs[pkey] = packs
	} // for

	if err = util.JsonWrite(content.JsonFilePath(), objs, content.Bom); err != nil {
		return fmt.Errorf("%s: write json failed: %w", content.TargetName(), err)
	} // if

	return nil
}
