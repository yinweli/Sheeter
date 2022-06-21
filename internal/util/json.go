package util

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

type JsonObj = map[string]interface{} // json物件型態
type JsonObjs = map[string]JsonObj    // json列表型態

// JsonWrite 寫入json檔案, 如果有需要會建立目錄
func JsonWrite(value any, filePath string, bom bool) error {
	bytes, err := json.MarshalIndent(value, "", "    ")

	if err != nil {
		return err
	} // if

	err = os.MkdirAll(path.Dir(filePath), os.ModePerm)

	if err != nil {
		return err
	} // if

	if bom {
		bytes = append([]byte{0xEF, 0xBB, 0xBF}[:], bytes[:]...)
	} // if

	err = ioutil.WriteFile(filePath, bytes, fs.ModePerm)

	if err != nil {
		return err
	} // if

	return nil
}
