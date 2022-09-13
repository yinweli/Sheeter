package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/utils"
)

// SectorJson 輸出json
func SectorJson(sector *Sector) error {
	rows, err := sector.GetRows(sector.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: sector json failed, data line not found", sector.StructName())
	} // if

	defer func() { _ = rows.Close() }()
	objs := map[string]interface{}{}

	for ok := true; ok; ok = rows.Next() {
		datas, _ := rows.Columns()

		if datas == nil {
			break // 碰到空行就結束了
		} // if

		packs, pkey, err := sector.layoutJson.Pack(datas, false)

		if err != nil {
			return fmt.Errorf("%s: sector json failed: %w", sector.StructName(), err)
		} // if

		objs[pkey] = packs
	} // for

	if err = utils.WriteJson(sector.FileJson(), objs); err != nil {
		return fmt.Errorf("%s: sector json failed: %w", sector.StructName(), err)
	} // if

	return nil
}

// SectorJsonSchema 輸出json架構
func SectorJsonSchema(sector *Sector) error {
	packs, _, err := sector.layoutJson.Pack([]string{}, true)

	if err != nil {
		return fmt.Errorf("%s: sector json schema failed: %w", sector.StructName(), err)
	} // if

	if err = utils.WriteJson(sector.FileJsonSchema(), packs); err != nil {
		return fmt.Errorf("%s: sector json schema failed: %w", sector.StructName(), err)
	} // if

	return nil
}
