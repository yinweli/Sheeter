package core

import (
	"bytes"
	"text/template"

	"github.com/ahmetb/go-linq/v3"
)

// CodeGenerate 產生程式碼
func CodeGenerate(code string, cargo *Cargo) (results []byte, err error) {
	temp, err := template.New("codeGenerate").Funcs(template.FuncMap{
		"cppNamespace": cppNamespace,
		"csNameSpace":  csNameSpace,
		"goPackage":    goPackage,
		"setline":      setline,
		"newline":      newline,
	}).Parse(code)

	if err != nil {
		return nil, err
	} // if

	buffer := &bytes.Buffer{}
	err = temp.Execute(buffer, cargo)

	if err != nil {
		return nil, err
	} // if

	return buffer.Bytes(), nil
}

// cppNamespace 取得c++命名空間名稱
func cppNamespace() string {
	return CppNamespace
}

// csNameSpace 取得c#命名空間名稱
func csNameSpace() string {
	return CsNamespace
}

// goPackage 取得go包名
func goPackage() string {
	return GoPackage
}

// setline 重設行數
func setline(columns []*Column) string {
	maxline = linq.From(columns).Where(func(itor interface{}) bool {
		return itor.(*Column).Field.Show()
	}).Count() - 1
	curline = 0
	return ""
}

// newline 取得新行
func newline() string {
	result := ""

	if maxline > curline {
		result = "\n"
	} // if

	curline++
	return result
}

var maxline = 0 // 最大行數
var curline = 0 // 當前行數
