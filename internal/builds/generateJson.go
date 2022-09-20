package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateJsonCsStruct 產生json-cs結構程式碼
func generateJsonCsStruct(runtimeStruct *RuntimeStruct) error {
	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.FileJsonCsStruct(), tmpls.JsonCsStruct.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate json-cs struct failed: %w", structName, err)
	} // if

	return nil
}

// generateJsonCsReader 產生json-cs讀取器程式碼
func generateJsonCsReader(runtimeStruct *RuntimeStruct) error {
	if runtimeStruct.Reader == false {
		return nil
	} // if

	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.FileJsonCsReader(), tmpls.JsonCsReader.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate json-cs reader failed: %w", structName, err)
	} // if

	return nil
}

// generateJsonGoStruct 產生json-go結構程式碼
func generateJsonGoStruct(runtimeStruct *RuntimeStruct) error {
	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.FileJsonGoStruct(), tmpls.JsonGoStruct.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate json-go struct failed: %w", structName, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", runtimeStruct.FileJsonGoStruct()); err != nil {
		return fmt.Errorf("%s: generate json-go struct failed, gofmt error: %w", structName, err)
	} // if

	return nil
}

// generateJsonGoReader 產生json-go讀取器程式碼
func generateJsonGoReader(runtimeStruct *RuntimeStruct) error {
	if runtimeStruct.Reader == false {
		return nil
	} // if

	structName := runtimeStruct.StructName()

	if err := utils.WriteTmpl(runtimeStruct.FileJsonGoReader(), tmpls.JsonGoReader.Data, runtimeStruct); err != nil {
		return fmt.Errorf("%s: generate json-go reader failed: %w", structName, err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", runtimeStruct.FileJsonGoReader()); err != nil {
		return fmt.Errorf("%s: generate json-go reader failed, gofmt error: %w", structName, err)
	} // if

	return nil
}
