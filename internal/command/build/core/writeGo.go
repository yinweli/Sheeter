package core

import (
	"fmt"
	"os/exec"

	"Sheeter/internal/util"
)

// WriteGo 寫入go
type WriteGo struct {
}

// Long 取得長名稱
func (this *WriteGo) Long() string {
	return "go"
}

// Short 取得短名稱
func (this *WriteGo) Short() string {
	return "g"
}

// Note 取得註解
func (this *WriteGo) Note() string {
	return "generate go file"
}

// Calc 計算進度值
func (this *WriteGo) Calc(sheetSize int) int {
	return 3
}

// Execute 執行工作
func (this *WriteGo) Execute(cargo *Cargo) (filePath string, err error) {
	cargo.Progress.Add(1)
	bytes, err := NewCoder(codeGo, cargo).Execute()

	if err != nil {
		return "", fmt.Errorf("convert go failed: %s [%s]", cargo.LogName(), err)
	} // if

	cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathGo, cargo.GoFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to go failed: %s [%s]", cargo.LogName(), err)
	} // if

	cargo.Progress.Add(1)
	cmd := exec.Command("go", "fmt", filePath)
	err = cmd.Run()

	if err != nil {
		return "", fmt.Errorf("write to go failed: %s [%s]", cargo.LogName(), err)
	} // if

	return filePath, nil
}

// codeGo go程式碼模板
var codeGo string = `// generation by sheeter ^o<

package {{goPackage}}

const {{.StructName}}FileName = "{{.JsonFileName}}" // json file name

type {{.StructName}} struct { {{setline .Columns}}
{{range .Columns}}{{if .Field.IsShow}}    {{.MemberName}} {{.Field.TypeGo}} // {{.Note}}{{newline}}{{end}}{{end}}
}`
