package core

import (
	"bytes"
	"text/template"
)

// Coder 程式碼編碼器
type Coder struct {
	Template string // 程式碼樣板
	Cargo    *Cargo // 工作箱
	maxline  int    // 最大行數
	curline  int    // 當前行數
}

// Execute 執行程式碼編碼
func (this *Coder) Execute() (results []byte, err error) {
	temp, err := template.New("coder").Funcs(template.FuncMap{
		"cppNamespace": this.cppNamespace,
		"csNamespace":  this.csNamespace,
		"goPackage":    this.goPackage,
		"setline":      this.setline,
		"newline":      this.newline,
	}).Parse(this.Template)

	if err != nil {
		return nil, err
	} // if

	buffer := &bytes.Buffer{}
	err = temp.Execute(buffer, this.Cargo)

	if err != nil {
		return nil, err
	} // if

	return buffer.Bytes(), nil
}

// cppNamespace 取得c++命名空間名稱
func (this *Coder) cppNamespace() string {
	return CppNamespace
}

// csNamespace 取得c#命名空間名稱
func (this *Coder) csNamespace() string {
	return CsNamespace
}

// goPackage 取得go包名
func (this *Coder) goPackage() string {
	return GoPackage
}

// setline 重設行數
func (this *Coder) setline(columns []*Column) string {
	this.maxline = 0

	for _, itor := range columns {
		if itor.Field.IsShow() {
			this.maxline++
		} // if
	} // for

	this.maxline = this.maxline - 1 // 減一是為了避免多換一次新行
	this.curline = 0
	return ""
}

// newline 取得新行
func (this *Coder) newline() string {
	result := ""

	if this.maxline > this.curline {
		result = "\n"
	} // if

	this.curline++
	return result
}

// NewCoder 建立程式碼編碼器
func NewCoder(code string, cargo *Cargo) *Coder {
	return &Coder{
		Template: code,
		Cargo:    cargo,
	}
}
