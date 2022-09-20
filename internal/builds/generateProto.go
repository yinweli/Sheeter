package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateProtoSchema 產生proto架構檔案
func generateProtoSchema(runtimeStruct *RuntimeStruct) error {
	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.FileProtoSchema(), tmpls.ProtoSchema.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate proto schema failed: %w", structName, err)
	} // if

	return nil
}
