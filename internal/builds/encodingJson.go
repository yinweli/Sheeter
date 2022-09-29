package builds

import (
	"fmt"

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

	json, err := layouts.JsonPack(rows, runtimeSector.layoutJson, runtimeSector.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(runtimeSector.PathJsonData(), json); err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	return nil
}
