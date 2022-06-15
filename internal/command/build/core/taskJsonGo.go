package core

import (
	"fmt"
	"os/exec"

	"github.com/yinweli/Sheeter/internal/util"
)

// jsonGoCode json/go程式碼模板
const jsonGoCode = `// generation by sheeter ^o<

package {{$.GoPackage}}

const {{$.StructName}}FileName = "{{$.JsonFileName}}" // json file name

type {{$.StructName}} struct { {{$.SetLine}}
{{range .Columns}}{{if .Field.IsShow}}    {{$.FirstUpper .Name}} {{.Field.TypeGo}} ` + "`json:\"{{.Name}}\"`" + ` // {{.Note}}{{$.NewLine}}{{end}}{{end}}
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

	err = exec.Command("go", "fmt", this.jsonGoFilePath()).Run()

	if err != nil {
		return fmt.Errorf("format go failed: %s [%s]", this.logName(), err)
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressM)
	} // if

	return nil
}
