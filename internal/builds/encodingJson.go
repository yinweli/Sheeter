package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingJson 產生json資料
func encodingJson(runtimeSector *RuntimeSector) error {
	structName := runtimeSector.StructName()
	line, err := runtimeSector.GetExcelLine(runtimeSector.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: data line not found", structName)
	} // if

	json, err := layouts.JsonPack(line, runtimeSector.layoutJson, runtimeSector.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(runtimeSector.JsonDataPath(), json); err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	return nil
}
