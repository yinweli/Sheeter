package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateProtoSchema 產生proto架構檔案
func generateProtoSchema(runtimeStruct *RuntimeStruct) error {
	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.PathProtoName(), tmpls.ProtoSchema.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate proto schema failed: %w", structName, err)
	} // if

	return nil
}

// generateProtoCsReader 產生json-cs讀取器程式碼
func generateProtoCsReader(runtimeStruct *RuntimeStruct) error {
	if runtimeStruct.Reader == false {
		return nil
	} // if

	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.PathProtoCsReader(), tmpls.ProtoCsReader.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate proto-cs reader failed: %w", structName, err)
	} // if

	return nil
}

// TODO: 還是要做protoReader喔(go)
