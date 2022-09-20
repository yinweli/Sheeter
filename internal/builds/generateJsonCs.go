package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateJsonCsStruct 產生json-cs結構程式碼
func generateJsonCsStruct(runtimeStruct *RuntimeStruct) error {
	if err := utils.WriteTmpl(runtimeStruct.Named.FileJsonCsStruct(), tmpls.JsonCsStruct.Code, runtimeStruct); err != nil {
		return fmt.Errorf("generate json-cs struct failed: %w", err)
	} // if

	return nil
}

// generateJsonCsReader 產生json-cs讀取器程式碼
func generateJsonCsReader(runtimeStruct *RuntimeStruct) error {
	if runtimeStruct.Reader {
		if err := utils.WriteTmpl(runtimeStruct.Named.FileJsonCsReader(), tmpls.JsonCsReader.Code, runtimeStruct); err != nil {
			return fmt.Errorf("generate json-cs reader failed: %w", err)
		} // if
	} // if

	return nil
}
