package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateProtoSchema 產生proto架構檔案
func generateProtoSchema(data *generateData) error {
	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoPath(), tmpls.ProtoSchema.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto schema failed: %w", structName, err)
	} // if

	return nil
}

// generateProtoCsReader 產生proto-cs讀取器程式碼
func generateProtoCsReader(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoCsReaderPath(), tmpls.ProtoCsReader.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto-cs reader failed: %w", structName, err)
	} // if

	return nil
}

// generateProtoGoReader 產生proto-go讀取器程式碼
func generateProtoGoReader(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoGoReaderPath(), tmpls.ProtoGoReader.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto-cs reader failed: %w", structName, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", data.ProtoGoReaderPath()); err != nil {
		return fmt.Errorf("%s: generate proto-go reader failed: gofmt error: %w", structName, err)
	} // if

	return nil
}
