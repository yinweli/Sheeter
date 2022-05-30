package core

import (
	"bytes"
	"text/template"

	"Sheeter/internal/util"
)

// CodeGenerate 產生程式碼
func CodeGenerate(code string, cargo *Cargo) (result string, err error) {
	temp, err := template.New("codeGenerate").Funcs(template.FuncMap{
		"excelName":  util.FirstUpper, // excel名稱為首字母大寫
		"sheetName":  util.FirstUpper, // sheet名稱為首字母大寫
		"memberName": util.FirstUpper, // 成員名稱為首字母大寫
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
