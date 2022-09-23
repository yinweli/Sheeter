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
	datas, err := JsonMarshal(value)

	if err != nil {
		return fmt.Errorf("json write failed: %w", err)
	} // if

	if err = WriteFile(path, datas); err != nil {
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

	if err = tmpl.Execute(buffer, refer); err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	if err = WriteFile(path, buffer.Bytes()); err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	return nil
}

// JsonMarshal 把物件轉換為json字串
func JsonMarshal(value any) (results []byte, err error) {
	return json.MarshalIndent(value, jsonPrefix, jsonIdent)
}
