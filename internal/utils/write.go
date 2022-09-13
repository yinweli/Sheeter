package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const jsonPrefix = ""    // json前綴字串
const jsonIdent = "    " // json縮排字串

// FileName 取得檔案名稱
func FileName(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// ExistFile 檔案是否存在
func ExistFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// WriteFile 寫入檔案, 如果有需要會建立目錄
func WriteFile(path string, datas []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	if err := os.WriteFile(path, datas, fs.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	return nil
}

// WriteJson 寫入json檔案, 如果有需要會建立目錄
func WriteJson(path string, value any) error {
	datas, err := json.MarshalIndent(value, jsonPrefix, jsonIdent)

	if err != nil {
		return fmt.Errorf("json write failed: %w", err)
	} // if

	err = WriteFile(path, datas)

	if err != nil {
		return fmt.Errorf("json write failed: %w", err)
	} // if

	return nil
}

// WriteTmpl 寫入模板檔案, 如果有需要會建立目錄
func WriteTmpl(path, content string, refer any) error {
	tmpl, err := template.New(path).Parse(content)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, refer)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	err = WriteFile(path, buffer.Bytes())

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	return nil
}
