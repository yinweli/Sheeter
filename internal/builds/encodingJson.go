package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingJson 產生json編碼資料
func encodingJson(runtimeSector *RuntimeSector) error {
	structName := runtimeSector.StructName()
	rows, err := runtimeSector.GetRows(runtimeSector.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed, data line not found", structName)
	} // if

	defer func() { _ = rows.Close() }()
	objs := map[string]interface{}{}

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		packs, pkey, err := runtimeSector.layoutJson.Pack(datas, false)

		if err != nil {
			return fmt.Errorf("%s: encoding json failed: %w", structName, err)
		} // if

		objs[pkey] = packs
	} // for

	if err = utils.WriteJson(runtimeSector.FileJsonDataPath(), objs); err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	return nil
}
