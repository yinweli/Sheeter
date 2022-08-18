package util

import (
	"bytes"
	"fmt"
	"text/template"
)

// TmplWrite 寫入模板檔案, 如果有需要會建立目錄
func TmplWrite(filePath, content string, refer any, bom bool) error {
	tmpl, err := template.New(filePath).Parse(content)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, refer)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	err = FileWrite(filePath, buffer.Bytes(), bom)

	if err != nil {
		return fmt.Errorf("tmpl write failed: %w", err)
	} // if

	return nil
}
