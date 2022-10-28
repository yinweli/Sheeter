package builds

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal"
)

// InitializeFile 初始化檔案
func InitializeFile(material any, result chan any) error {
	data, ok := material.(string)

	if ok == false {
		return nil
	} // if

	err := filepath.Walk(data, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		} // if

		if info.IsDir() {
			return nil
		} // if

		if filepath.Ext(path) != internal.ExcelExt {
			return nil
		} // if

		result <- path
		return nil
	})

	if err != nil {
		return fmt.Errorf("%s: initialize path failed: %w", data, err)
	} // if

	return nil
}
