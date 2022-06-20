package util

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

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

// JsonRead 讀取json檔案
func JsonRead(value any, filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	} // if

	err = json.Unmarshal(bytes, value)

	if err != nil {
		return err
	} // if

	return nil
}
