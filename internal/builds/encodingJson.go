package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingJson 產生json資料
func encodingJson(data *encodingData) error {
	structName := data.StructName()
	line, err := data.excel.GetLine(data.Sheet, data.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: data line not found", structName)
	} // if

	json, err := layouts.JsonPack(line, data.layoutJson, data.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(data.JsonDataPath(), json); err != nil {
		return fmt.Errorf("%s: encoding json failed: %w", structName, err)
	} // if

	return nil
}
