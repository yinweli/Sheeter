package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateProtoSchema 產生proto架構檔案
func generateProtoSchema(runtimeStruct *RuntimeStruct) error {
	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.ProtoPath(), tmpls.ProtoSchema.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate proto schema failed: %w", structName, err)
	} // if

	return nil
}

// generateProtoCsReader 產生proto-cs讀取器程式碼
func generateProtoCsReader(runtimeStruct *RuntimeStruct) error {
	if runtimeStruct.Reader == false {
		return nil
	} // if

	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.ProtoCsReaderPath(), tmpls.ProtoCsReader.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate proto-cs reader failed: %w", structName, err)
	} // if

	return nil
}

// generateProtoGoReader 產生proto-go讀取器程式碼
func generateProtoGoReader(runtimeStruct *RuntimeStruct) error {
	if runtimeStruct.Reader == false {
		return nil
	} // if

	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.ProtoGoReaderPath(), tmpls.ProtoGoReader.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate proto-cs reader failed: %w", structName, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", runtimeStruct.ProtoGoReaderPath()); err != nil {
		return fmt.Errorf("%s: generate proto-go reader failed: gofmt error: %w", structName, err)
	} // if

	return nil
}
