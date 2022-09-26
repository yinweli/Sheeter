package builds

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingJson 產生json資料
func encodingJson(runtimeSector *RuntimeSector) error {
	structName := runtimeSector.StructName()
	rows, err := runtimeSector.GetRows(runtimeSector.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: data line not found", structName)
	} // if

	json, err := packJson(rows, runtimeSector.layoutJson)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(runtimeSector.FileJsonDataPath(), json); err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	return nil
}

// packJson
func packJson(rows *excelize.Rows, layoutJson *layouts.LayoutJson) (json []byte, err error) {
	defer func() { _ = rows.Close() }()
	objs := map[int64]interface{}{}

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		packs, pkey, err := layoutJson.Pack(datas, false)

		if err != nil {
			return nil, fmt.Errorf("pack json failed: %w", err)
		} // if

		objs[pkey] = packs
	} // for

	if json, err = utils.JsonMarshal(objs); err != nil {
		return nil, fmt.Errorf("pack json failed: %w", err)
	} // if

	return json, nil
}
