package core

import (
	"fmt"

	"Sheeter/internal/util"
)

// WriteCs 寫入c#
func WriteCs(cargo *Cargo) (filePath string, err error) {
	bytes, err := CodeGenerate(codeCs, cargo)

	if err != nil {
		return "", fmt.Errorf("convert cs failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	_ = cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathCs, cargo.CsFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to cs failed: %s [%s]", cargo.Element.GetFullName(), err)
	} // if

	return filePath, nil
}

// codeCs c#程式碼模板
var codeCs string = `// generation by sheeter ^o<

using System;
using System.Collections.Generic;

namespace {{csNamespace}} {
    public class {{.StructName}} { {{setline .Columns}}
        public const string fileName = "{{.JsonFileName}}";
{{range .Columns}}{{if .Field.Show}}        public {{.Field.TypeCs}} {{.MemberName}}; // {{.Note}}{{newline}}{{end}}{{end}}
    }
} // namespace {{csNamespace}}`
