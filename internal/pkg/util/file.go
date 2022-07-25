package util

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var bomPrefix = []byte{0xEF, 0xBB, 0xBF} // bom前置資料

// FileWrite 寫入檔案, 如果有需要會建立目錄
func FileWrite(filePath string, bytes []byte, bom bool) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	if bom {
		bytes = append(bomPrefix, bytes...)
	} // if

	if err := os.WriteFile(filePath, bytes, fs.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	return nil
}
