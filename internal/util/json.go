package util

import (
	"encoding/json"
)

const jsonPrefix = ""    // json前綴字串
const jsonIdent = "    " // json縮排字串

type JsonObj = map[string]interface{} // json物件型態
type JsonObjs = map[string]JsonObj    // json列表型態

// JsonWrite 寫入json檔案, 如果有需要會建立目錄
func JsonWrite(filePath string, value any, bom bool) error {
	bytes, err := json.MarshalIndent(value, jsonPrefix, jsonIdent)

	if err != nil {
		return err
	} // if

	err = FileWrite(filePath, bytes, bom)

	if err != nil {
		return err
	} // if

	return nil
}
