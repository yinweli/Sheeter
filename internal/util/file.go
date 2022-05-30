package util

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFile 寫入檔案
func WriteFile(path string, name string, bytes []byte) error {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	} // if

	err = ioutil.WriteFile(filepath.Join(path, name), bytes, fs.ModePerm)

	if err != nil {
		return err
	} // if

	return nil
}
