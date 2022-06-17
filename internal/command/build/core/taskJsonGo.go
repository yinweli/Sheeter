package core

import (
	"fmt"
	"os/exec"

	"github.com/yinweli/Sheeter/internal/util"
)

// jsonGoCode json/go程式碼模板
const jsonGoCode = `// generation by sheeter ^o<, from {{$.OriginalName}}
package {{$.GoPackage}}

import "encoding/json"

const {{$.StructName}}FileName = "{{$.JsonFileName}}" // json file name

type {{$.StructName}} struct { {{$.SetLine}}
{{range .Columns}}{{if .Field.IsShow}}    {{$.FirstUpper .Name}} {{.Field.TypeGo}} ` + "`json:\"{{.Name}}\"`" + ` // {{.Note}}{{$.NewLine}}{{end}}{{end}}
}

type {{$.StructName}}Map map[int]{{$.StructName}}

func (this *{{$.StructName}}Map) ParseString(s string) error {
    return json.Unmarshal([]byte(s), this)
}

func (this *{{$.StructName}}Map) ParseBytes(b []byte) error {
    return json.Unmarshal(b, this)
}
`

// executeJsonGo 輸出json/go
func (this *Task) executeJsonGo() error {
	stemplateCode := STemplateCode{
		STemplate: STemplate{
			OriginalName: this.originalName(),
			StructName:   this.structName(),
		},
		JsonFileName: this.jsonFileName(),
		Columns:      this.columns,
	}
	bytes, err := stemplateCode.Generate(jsonGoCode)

	if err != nil {
		return fmt.Errorf("generate go failed: %s [%s]", this.originalName(), err)
	} // if

	err = util.FileWrite(this.jsonGoFilePath(), bytes, this.global.Bom)

	if err != nil {
		return fmt.Errorf("write to go failed: %s [%s]", this.originalName(), err)
	} // if

	err = exec.Command("go", "fmt", this.jsonGoFilePath()).Run()

	if err != nil {
		return fmt.Errorf("format go failed: %s [%s]", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressM)
	} // if

	return nil
}
