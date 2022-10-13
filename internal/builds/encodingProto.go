package builds

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingProto 產生proto資料
func encodingProto(data *encodingData) error {
	structName := data.StructName()
	sheet, err := data.excel.Get(data.Sheet)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: sheet not found", structName)
	} // if

	sheet.Nextn(data.LineOfData)
	json, err := layouts.JsonPack(sheet, data.layoutJson, data.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	filename := data.ProtoName()
	message := data.StorerMessage(data.SimpleNamespace)
	imports := []string{filepath.Join(internal.ProtoPath, internal.SchemaPath)}
	proto, err := utils.JsonToProto(filename, message, imports, json)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(data.ProtoDataPath(), proto); err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	return nil
}
