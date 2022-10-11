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
	line, err := data.excel.GetLine(data.Sheet, data.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: data line not found", structName)
	} // if

	json, err := layouts.JsonPack(line, data.layoutJson, data.Excludes)

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
