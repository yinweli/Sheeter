package util

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

// FileWrite 寫入檔案, 如果有需要會建立目錄
func FileWrite(filePath string, bytes []byte, bom bool) error {
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)

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

// FileSize 取得檔案長度
func FileSize(file *os.File) int64 {
	stat, err := file.Stat()

	if err != nil {
		return 0
	} // if

	return stat.Size()
}
