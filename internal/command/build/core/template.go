package core

import "github.com/yinweli/Sheeter/internal/util"

// Template 產生器樣板
type Template struct {
	OriginalName string // 原始名稱
	StructName   string // 結構名稱
	maxline      int    // 最大行數
	curline      int    // 當前行數
}

// SetLine 設置行數
func (this *Template) SetLine(maxline int) string {
	this.maxline = maxline
	this.curline = 0
	return ""
}

// NewLine 換行輸出
func (this *Template) NewLine() string {
	defer func() { this.curline++ }()

	if this.maxline > this.curline {
		return "\n"
	} // if

	return ""
}

// FirstUpper 首字大寫
func (this *Template) FirstUpper(input string) string {
	return util.FirstUpper(input)
}

// FirstLower 首字小寫
func (this *Template) FirstLower(input string) string {
	return util.FirstLower(input)
}

// NewTemplate 建立產生器樣板
func NewTemplate(originalName string, structName string) Template {
	return Template{
		OriginalName: originalName,
		StructName:   structName,
	}
}
