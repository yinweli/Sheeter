package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// FileName 取得檔案名稱
func FileName(path string) string {
	path = filepath.Clean(path) // 避免奇怪的相對路徑
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// FileExist 檔案是否存在
func FileExist(path string) bool {
	path = filepath.Clean(path) // 避免奇怪的相對路徑
	stat, err := os.Stat(path)
	return os.IsNotExist(err) == false && stat != nil && stat.IsDir() == false
}

// WriteFile 寫入檔案, 如果有需要會建立目錄
func WriteFile(path string, data []byte) error {
	const dirPerm = 0750  // 目錄權限
	const filePerm = 0644 // 檔案權限

	path = filepath.Clean(path) // 避免奇怪的相對路徑

	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("writeFile: %w", err)
	} // if

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("writeFile: %w", err)
	} // if

	return nil
}

// WriteTmpl 寫入模板檔案, 如果有需要會建立目錄
func WriteTmpl(path, content string, refer any) error {
	path = filepath.Clean(path) // 避免奇怪的相對路徑
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
