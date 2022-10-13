package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingJson 產生json資料
func encodingJson(data *encodingData) error {
	structName := data.StructName()
	sheet, err := data.excel.Get(data.Sheet)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: sheet not found", structName)
	} // if

	sheet.Nextn(data.LineOfData)
	json, err := layouts.JsonPack(sheet, data.layoutJson, data.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(data.JsonDataPath(), json); err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	return nil
}
