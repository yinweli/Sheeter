package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateJsonCsStruct 產生json-cs結構程式碼
func generateJsonCsStruct(data *generateData) error {
	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonCsStructPath(), tmpls.JsonCsStruct.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-cs struct failed: %w", structName, err)
	} // if

	return nil
}

// generateJsonCsReader 產生json-cs讀取器程式碼
func generateJsonCsReader(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonCsReaderPath(), tmpls.JsonCsReader.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-cs reader failed: %w", structName, err)
	} // if

	return nil
}

// generateJsonGoStruct 產生json-go結構程式碼
func generateJsonGoStruct(data *generateData) error {
	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonGoStructPath(), tmpls.JsonGoStruct.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-go struct failed: %w", structName, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", data.JsonGoStructPath()); err != nil {
		return fmt.Errorf("%s: generate json-go struct failed: gofmt error: %w", structName, err)
	} // if

	return nil
}

// generateJsonGoReader 產生json-go讀取器程式碼
func generateJsonGoReader(data *generateData) error {
	if data.Reader == false {
		return nil
	} // if

	structName := data.StructName()

	if err := utils.WriteTmpl(data.JsonGoReaderPath(), tmpls.JsonGoReader.Data, data); err != nil {
		return fmt.Errorf("%s: generate json-go reader failed: %w", structName, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", data.JsonGoReaderPath()); err != nil {
		return fmt.Errorf("%s: generate json-go reader failed: gofmt error: %w", structName, err)
	} // if

	return nil
}
