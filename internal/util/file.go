package util

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileWrite 寫入檔案
func FileWrite(path string, name string, bytes []byte) (filePath string, err error) {
	err = os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return "", err
	} // if

	filePath = filepath.Join(path, name)
	err = ioutil.WriteFile(filePath, bytes, fs.ModePerm)

	if err != nil {
		return "", err
	} // if

	return filePath, nil
}
