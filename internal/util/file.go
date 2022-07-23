package util

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

var bomPrefix = []byte{0xEF, 0xBB, 0xBF} // bom前置資料

// FileWrite 寫入檔案, 如果有需要會建立目錄
func FileWrite(filePath string, bytes []byte, bom bool) error {
	err := os.MkdirAll(path.Dir(filePath), os.ModePerm)

	if err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	if bom {
		bytes = append(bomPrefix, bytes...)
	} // if

	err = os.WriteFile(filePath, bytes, fs.ModePerm)

	if err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	return nil
}
