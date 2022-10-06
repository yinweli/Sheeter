package builds

import (
	"fmt"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/layouts"
	"github.com/yinweli/Sheeter/internal/utils"
)

// encodingProto 產生proto資料
func encodingProto(runtimeSector *RuntimeSector) error {
	structName := runtimeSector.StructName()
	rows, err := runtimeSector.GetRows(runtimeSector.LineOfData)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: data line not found", structName)
	} // if

	json, err := layouts.JsonPack(rows, runtimeSector.layoutJson, runtimeSector.Excludes)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	filename := runtimeSector.ProtoName()
	message := runtimeSector.StorerMessage(runtimeSector.Global.SimpleNamespace)
	imports := []string{filepath.Join(internal.ProtoPath, internal.SchemaPath)}
	data, err := utils.JsonToProto(filename, message, imports, json)

	if err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	if err := utils.WriteFile(runtimeSector.ProtoDataPath(), data); err != nil {
		return fmt.Errorf("%s: encoding proto failed: %w", structName, err)
	} // if

	return nil
}
