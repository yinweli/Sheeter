package core

import (
	"fmt"
	"os/exec"

	"Sheeter/internal/util"
)

// WriteGo 寫入go
func WriteGo(cargo *Cargo) (filePath string, err error) {
	bytes, err := CodeGenerate(codeGo, cargo)

	if err != nil {
		return "", fmt.Errorf("convert go failed: %s [%s]", cargo.LogName(), err)
	} // if

	cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathGo, cargo.GoFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to go failed: %s [%s]", cargo.LogName(), err)
	} // if

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

const {{.StructName}}FileName string = "{{.JsonFileName}}" // json file name

type {{.StructName}} struct { {{setline .Columns}}
{{range .Columns}}{{if .Field.Show}}    {{.MemberName}} {{.Field.TypeGo}} // {{.Note}}{{newline}}{{end}}{{end}}
}`
