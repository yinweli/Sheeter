package core

import (
	"bytes"
	"text/template"

	"Sheeter/internal/util"

	"github.com/ahmetb/go-linq/v3"
)

// CodeGenerate 產生程式碼
func CodeGenerate(code string, cargo *Cargo) (results []byte, err error) {
	temp, err := template.New("codeGenerate").Funcs(template.FuncMap{
		"cppNamespace": cppNamespace,
		"csNameSpace":  csNameSpace,
		"goPackage":    goPackage,
		"structName":   structName,
		"memberName":   memberName,
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

// structName 取得結構名稱
func structName(cargo *Cargo) string {
	return util.FirstUpper(cargo.PureExcelName()) + util.FirstUpper(cargo.Element.Sheet)
}

// memberName 取得成員名稱
func memberName(name string) string {
	return util.FirstUpper(name)
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
