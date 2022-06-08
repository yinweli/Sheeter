package core

import (
	"fmt"

	"Sheeter/internal/util"
)

// jsonCsCode json/c#程式碼模板
const jsonCsCode = `// generation by sheeter ^o<

using System;
using System.Collections.Generic;

namespace {{.CsNamespace}} {
    public class {{.StructName}} { {{setline .Columns}}
        public const string fileName = "{{.JsonFileName}}";
{{range .Columns}}{{if .Field.IsShow}}        public {{.Field.TypeCs}} {{.ColumnName}}; // {{.Note}}{{newline}}{{end}}{{end}}
    }
} // namespace {{.CsNamespace}}
`

// TaskJsonCs 輸出json/cs
func TaskJsonCs(ctx *Context) error {
	bytes, err := Coder(jsonCsCode, ctx)

	if err != nil {
		return fmt.Errorf("generate cs failed: %s [%s]", ctx.LogName(), err)
	} // if

	err = util.FileWrite(ctx.JsonCsFilePath(), bytes)

	if err != nil {
		return fmt.Errorf("write to cs failed: %s [%s]", ctx.LogName(), err)
	} // if

	_ = ctx.Progress.Add(1)
	return nil
}
