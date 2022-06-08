package core

import (
	"bytes"
	"text/template"
)

// Coder 執行程式碼編碼
func Coder(code string, ctx *Context) (results []byte, err error) {
	maxline := 0
	curline := 0
	temp, err := template.New("coder").Funcs(template.FuncMap{
		"setline": func(columns []*Column) string {
			maxline = 0

			for _, itor := range columns {
				if itor.Field.IsShow() {
					maxline++
				} // if
			} // for

			maxline = maxline - 1 // 減一是為了避免多換一次新行
			curline = 0
			return ""
		},
		"newline": func() string {
			result := ""

			if maxline > curline {
				result = "\n"
			} // if

			curline++
			return result
		},
	}).Parse(code)

	if err != nil {
		return nil, err
	} // if

	buffer := &bytes.Buffer{}
	err = temp.Execute(buffer, ctx)

	if err != nil {
		return nil, err
	} // if

	return buffer.Bytes(), nil
}
