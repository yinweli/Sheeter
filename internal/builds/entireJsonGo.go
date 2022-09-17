package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/codes"
	"github.com/yinweli/Sheeter/internal/utils"
)

// EntireJsonGoStruct 輸出json-go程式碼
func EntireJsonGoStruct(entire *Entire) error {
	if err := utils.WriteTmpl(entire.Named.FileJsonGoStruct(), codes.JsonGoStruct.Code, entire); err != nil {
		return fmt.Errorf("entire json-go struct failed: %w", err)
	} // if

	if err := utils.ShellRun("gofmt", "-w", entire.Named.FileJsonGoStruct()); err != nil {
		return fmt.Errorf("sector json-go struct failed, gofmt error: %w", err)
	} // if

	return nil
}

// EntireJsonGoReader 輸出json-go讀取器
func EntireJsonGoReader(entire *Entire) error {
	if entire.Reader {
		if err := utils.WriteTmpl(entire.Named.FileJsonGoReader(), codes.JsonGoReader.Code, entire); err != nil {
			return fmt.Errorf("entire json-go reader failed: %w", err)
		} // if

		if err := utils.ShellRun("gofmt", "-w", entire.Named.FileJsonGoReader()); err != nil {
			return fmt.Errorf("sector json-go reader failed, gofmt error: %w", err)
		} // if
	} // if

	return nil
}
