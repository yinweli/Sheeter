package core

import "github.com/yinweli/Sheeter/internal/util"

// STemplate 產生器樣板
type STemplate struct {
	OriginalName string // 原始名稱
	StructName   string // 結構名稱
	maxline      int    // 最大行數
	curline      int    // 當前行數
}

// SetLine 設置行數
func (this *STemplate) SetLine(maxline int) string {
	this.maxline = maxline
	this.curline = 0
	return ""
}

// NewLine 換行輸出
func (this *STemplate) NewLine() string {
	defer func() { this.curline++ }()

	if this.maxline > this.curline {
		return "\n"
	} // if

	return ""
}

// FirstUpper 首字大寫
func (this *STemplate) FirstUpper(input string) string {
	return util.FirstUpper(input)
}

// FirstLower 首字小寫
func (this *STemplate) FirstLower(input string) string {
	return util.FirstLower(input)
}

func NewSTemplate(originalName string, structName string) *STemplate {
	return &STemplate{
		OriginalName: originalName,
		StructName:   structName,
	}
}
