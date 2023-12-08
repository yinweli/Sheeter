package utils

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// FileName 取得檔案名稱
func FileName(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// FileExist 檔案是否存在
func FileExist(path string) bool {
	stat, err := os.Stat(path)
	return os.IsNotExist(err) == false && stat != nil && stat.IsDir() == false
}

// WriteFile 寫入檔案, 如果有需要會建立目錄
func WriteFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return fmt.Errorf("writeFile: %w", err)
	} // if

	if err := os.WriteFile(path, data, fs.ModePerm); err != nil {
		return fmt.Errorf("writeFile: %w", err)
	} // if

	return nil
}

// WriteTmpl 寫入模板檔案, 如果有需要會建立目錄
func WriteTmpl(path, content string, refer any) error {
	tmpl, err := template.New(path).Parse(content)

	if err != nil {
		return fmt.Errorf("writeTmpl: %w", err)
	} // if

	buffer := &bytes.Buffer{}

	if err = tmpl.Execute(buffer, refer); err != nil {
		return fmt.Errorf("writeTmpl: %w", err)
	} // if

	if err = WriteFile(path, buffer.Bytes()); err != nil {
		return fmt.Errorf("writeTmpl: %w", err)
	} // if

	return nil
}
