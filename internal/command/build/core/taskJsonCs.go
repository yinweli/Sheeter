package core

import (
	"fmt"

	"Sheeter/internal/util"
)

// jsonCsCode json/c#程式碼模板
const jsonCsCode = `// generation by sheeter ^o<

using System;
using System.Collections.Generic;

namespace {{$.CsNamespace}} {
    public class {{$.StructName}} { {{$.SetLine}}
        public const string fileName = "{{$.JsonFileName}}";
{{range .Columns}}{{if .Field.IsShow}}        public {{.Field.TypeCs}} {{$.ColumnName .Name}}; // {{.Note}}{{$.NewLine}}{{end}}{{end}}
    }
} // namespace {{$.CsNamespace}}
`

// executeJsonCs 輸出json/cs
func (this *Task) executeJsonCs() error {
	bytes, err := NewCoder(this.columns, this.global.CppLibraryPath, this.jsonFileName(), this.structName()).Generate(jsonCsCode)

	if err != nil {
		return fmt.Errorf("generate cs failed: %s [%s]", this.logName(), err)
	} // if

	err = util.FileWrite(this.jsonCsFilePath(), bytes, this.global.Bom)

	if err != nil {
		return fmt.Errorf("write to cs failed: %s [%s]", this.logName(), err)
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressM)
	} // if

	return nil
}
