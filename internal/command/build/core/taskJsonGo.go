package core

import (
	"fmt"
	"os/exec"

	"Sheeter/internal/util"
)

// jsonGoCode json/go程式碼模板
const jsonGoCode = `// generation by sheeter ^o<

package {{$.GoPackage}}

const {{$.StructName}}FileName = "{{$.JsonFileName}}" // json file name

type {{$.StructName}} struct { {{$.SetLine}}
{{range .Columns}}{{if .Field.IsShow}}    {{$.ColumnName .Name}} {{.Field.TypeGo}} // {{.Note}}{{$.NewLine}}{{end}}{{end}}
}
`

// executeJsonGo 輸出json/go
func (this *Task) executeJsonGo() error {
	bytes, err := NewCoder(this.columns, this.global.CppLibraryPath, this.jsonFileName(), this.structName()).Generate(jsonGoCode)

	if err != nil {
		return fmt.Errorf("generate go failed: %s [%s]", this.logName(), err)
	} // if

	err = util.FileWrite(this.jsonGoFilePath(), bytes, this.global.Bom)

	if err != nil {
		return fmt.Errorf("write to go failed: %s [%s]", this.logName(), err)
	} // if

	cmd := exec.Command("go", "fmt", this.jsonGoFilePath())
	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("format go failed: %s [%s]", this.logName(), err)
	} // if

	return nil
}
