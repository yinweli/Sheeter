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

// generateProtoReaderCs 產生proto讀取器cs程式碼
func generateProtoReaderCs(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoReaderCsPath(), tmpls.ProtoReaderCs.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto-reader-cs failed: %w", structName, err)
	} // if

	return nil
}

// generateProtoReaderGo 產生proto-go讀取器程式碼
func generateProtoReaderGo(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.ProtoReaderGoPath(), tmpls.ProtoReaderGo.Data, data); err != nil {
		return fmt.Errorf("%s: generate proto-reader-cs failed: %w", structName, err)
	} // if

	return nil
}
