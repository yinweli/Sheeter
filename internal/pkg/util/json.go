package util

import (
	"encoding/json"
	"fmt"
)

const jsonPrefix = ""    // json前綴字串
const jsonIdent = "    " // json縮排字串

// JsonWrite 寫入json檔案, 如果有需要會建立目錄
func JsonWrite(filePath string, value any, bom bool) error {
	bytes, err := json.MarshalIndent(value, jsonPrefix, jsonIdent)

	if err != nil {
		return fmt.Errorf("json write failed: %w", err)
	} // if

	err = FileWrite(filePath, bytes, bom)

	if err != nil {
		return fmt.Errorf("json write failed: %w", err)
	} // if

	return nil
}
