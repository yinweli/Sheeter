package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateJsonStructCs 產生json結構cs程式碼
func generateJsonStructCs(data *generateData) error {
	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonStructCsPath(), tmpls.JsonStructCs.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-struct-cs failed: %w", structName, err)
	} // if

	return nil
}

// generateJsonReaderCs 產生json讀取器cs程式碼
func generateJsonReaderCs(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonReaderCsPath(), tmpls.JsonReaderCs.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-reader-cs failed: %w", structName, err)
	} // if

	return nil
}

// generateJsonStructGo 產生json結構go程式碼
func generateJsonStructGo(data *generateData) error {
	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonStructGoPath(), tmpls.JsonStructGo.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-struct-go failed: %w", structName, err)
	} // if

	return nil
}

// generateJsonReaderGo 產生json讀取器go程式碼
func generateJsonReaderGo(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonReaderGoPath(), tmpls.JsonReaderGo.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-reader-go failed: %w", structName, err)
	} // if

	return nil
}
