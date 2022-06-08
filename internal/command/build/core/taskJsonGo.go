package core

import (
	"fmt"
	"os/exec"

	"Sheeter/internal/util"
)

// jsonGoCode json/go程式碼模板
const jsonGoCode = `// generation by sheeter ^o<

package {{.GoPackage}}

const {{.StructName}}FileName = "{{.JsonFileName}}" // json file name

type {{.StructName}} struct { {{setline .Columns}}
{{range .Columns}}{{if .Field.IsShow}}    {{.ColumnName}} {{.Field.TypeGo}} // {{.Note}}{{newline}}{{end}}{{end}}
}
`

// TaskJsonGo 輸出json/go
func TaskJsonGo(ctx *Context) error {
	bytes, err := Coder(jsonGoCode, ctx)

	if err != nil {
		return fmt.Errorf("generate go failed: %s [%s]", ctx.LogName(), err)
	} // if

	err = util.FileWrite(ctx.JsonGoFilePath(), bytes)

	if err != nil {
		return fmt.Errorf("write to go failed: %s [%s]", ctx.LogName(), err)
	} // if

	cmd := exec.Command("go", "fmt", ctx.JsonGoFilePath())
	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("format go failed: %s [%s]", ctx.LogName(), err)
	} // if

	_ = ctx.Progress.Add(1)
	return nil
}
