package core

import (
	"bytes"
	"text/template"

	"Sheeter/internal/util"
)

// CodeGenerate 產生程式碼
func CodeGenerate(code string, cargo *Cargo) (result string, err error) {
	temp, err := template.New("codeGenerate").Funcs(template.FuncMap{
		"structName": structName,
		"memberName": memberName,
	}).Parse(code)

	if err != nil {
		return "", err
	} // if

	buffer := &bytes.Buffer{}
	err = temp.Execute(buffer, cargo)

	if err != nil {
		return "", err
	} // if

	return buffer.String(), nil
}

// structName 取得結構名稱
func structName(cargo *Cargo) string {
	return util.FirstUpper(cargo.Element.Excel) + util.FirstUpper(cargo.Element.Sheet)
}

// memberName 取得成員名稱
func memberName(name string) string {
	return util.FirstUpper(name)
}
