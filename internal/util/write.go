package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

const jsonPrefix = ""    // json前綴字串
const jsonIdent = "    " // json縮排字串

var bomPrefix = []byte{0xEF, 0xBB, 0xBF} // bom前置資料

// WriteFile 寫入檔案, 如果有需要會建立目錄
func WriteFile(filePath string, datas []byte, bom bool) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	if bom {
		datas = append(bomPrefix, datas...)
	} // if

	if err := os.WriteFile(filePath, datas, fs.ModePerm); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	} // if

	return nil
}

// WriteJson 寫入json檔案, 如果有需要會建立目錄
func WriteJson(filePath string, value any, bom bool) error {
	datas, err := json.MarshalIndent(value, jsonPrefix, jsonIdent)

	if err != nil {
		return fmt.Errorf("json write failed: %w", err)
	} // if

	err = WriteFile(filePath, datas, bom)

	if err != nil {
		return fmt.Errorf("json write failed: %w", err)
	} // if

	return nil
}

// WriteTmpl 寫入模板檔案, 如果有需要會建立目錄
func WriteTmpl(filePath, content string, refer any, bom bool) error {
	tmpl, err := template.New(filePath).Parse(content)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, refer)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	err = WriteFile(filePath, buffer.Bytes(), bom)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	return nil
}
