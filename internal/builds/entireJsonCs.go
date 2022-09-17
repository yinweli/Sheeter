package builds

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/codes"
	"github.com/yinweli/Sheeter/internal/utils"
)

// EntireJsonCsStruct 輸出json-cs程式碼
func EntireJsonCsStruct(entire *Entire) error {
	if err := utils.WriteTmpl(entire.Named.FileJsonCsStruct(), codes.JsonCsStruct.Code, entire); err != nil {
		return fmt.Errorf("entire json-cs struct failed: %w", err)
	} // if

	return nil
}

// EntireJsonCsReader 輸出json-cs讀取器
func EntireJsonCsReader(entire *Entire) error {
	if entire.Reader {
		if err := utils.WriteTmpl(entire.Named.FileJsonCsReader(), codes.JsonCsReader.Code, entire); err != nil {
			return fmt.Errorf("entire json-cs reader failed: %w", err)
		} // if
	} // if

	return nil
}
