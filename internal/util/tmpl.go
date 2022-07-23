package util

import (
	"bytes"
	"fmt"
	"text/template"
)

// TmplWrite 寫入模板檔案, 如果有需要會建立目錄
func TmplWrite(filePath string, bom bool, content string, refer any) error {
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

// TmplLine 換行模板工具, 模板資料包含此結構可以獲得換行功能
type TmplLine struct {
	maxline int // 最大行數
	curline int // 當前行數
}

// SetLine 設置行數
func (this *TmplLine) SetLine(maxline int) string {
	this.maxline = maxline
	this.curline = 0
	return ""
}

// NewLine 換行輸出
func (this *TmplLine) NewLine() string {
	defer func() { this.curline++ }()

	if this.maxline > this.curline {
		return "\n"
	} // if

	return ""
}

// TmplFirstChar 首字模板工具, 模板資料包含此結構可以獲得首字功能
type TmplFirstChar struct {
}

// FirstUpper 首字大寫
func (this *TmplFirstChar) FirstUpper(input string) string {
	return FirstUpper(input)
}

// FirstLower 首字小寫
func (this *TmplFirstChar) FirstLower(input string) string {
	return FirstLower(input)
}
