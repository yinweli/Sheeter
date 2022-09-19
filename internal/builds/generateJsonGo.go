package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/tmpls"
	"github.com/yinweli/Sheeter/internal/utils"
)

// generateJsonGoStruct 產生json-go結構程式碼
func generateJsonGoStruct(runtimeStruct *RuntimeStruct) error {
	if err := utils.WriteTmpl(runtimeStruct.Named.FileJsonGoStruct(), tmpls.JsonGoStruct.Code, runtimeStruct); err != nil {
		return fmt.Errorf("generate json-go struct failed: %w", err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", runtimeStruct.Named.FileJsonGoStruct()); err != nil {
		return fmt.Errorf("generate json-go struct failed, gofmt error: %w", err)
	} // if

	return nil
}

// generateJsonGoReader 產生json-go讀取器程式碼
func generateJsonGoReader(runtimeStruct *RuntimeStruct) error {
	if runtimeStruct.Reader {
		if err := utils.WriteTmpl(runtimeStruct.Named.FileJsonGoReader(), tmpls.JsonGoReader.Code, runtimeStruct); err != nil {
			return fmt.Errorf("generate json-go reader failed: %w", err)
		} // if

		if err := utils.ShellRun("gofmt", "-w", runtimeStruct.Named.FileJsonGoReader()); err != nil {
			return fmt.Errorf("generate json-go reader failed, gofmt error: %w", err)
		} // if
	} // if

	return nil
}
