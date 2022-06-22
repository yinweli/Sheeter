package util

// TempTool 模板產生工具
type TempTool struct {
	maxline int // 最大行數
	curline int // 當前行數
}

// SetLine 設置行數
func (this *TempTool) SetLine(maxline int) string {
	this.maxline = maxline
	this.curline = 0
	return ""
}

// NewLine 換行輸出
func (this *TempTool) NewLine() string {
	defer func() { this.curline++ }()

	if this.maxline > this.curline {
		return "\n"
	} // if

	return ""
}

// FirstUpper 首字大寫
func (this *TempTool) FirstUpper(input string) string {
	return FirstUpper(input)
}

// FirstLower 首字小寫
func (this *TempTool) FirstLower(input string) string {
	return FirstLower(input)
}
