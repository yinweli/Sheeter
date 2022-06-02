package core

import (
	"fmt"

	"Sheeter/internal/util"
)

// WriteCs 寫入c#
type WriteCs struct {
}

// LongName 取得長名稱
func (this *WriteCs) LongName() string {
	return "cs"
}

// ShortName 取得短名稱
func (this *WriteCs) ShortName() string {
	return "s"
}

// Note 取得註解
func (this *WriteCs) Note() string {
	return "generate cs file"
}

// Progress 取得進度值
func (this *WriteCs) Progress(sheetSize int) int {
	return 2
}

// Execute 執行工作
func (this *WriteCs) Execute(cargo *Cargo) (filePath string, err error) {
	cargo.Progress.Add(1)
	bytes, err := CodeGenerate(codeCs, cargo)

	if err != nil {
		return "", fmt.Errorf("convert cs failed: %s [%s]", cargo.LogName(), err)
	} // if

	cargo.Progress.Add(1)
	filePath, err = util.FileWrite(OutputPathCs, cargo.CsFileName(), bytes)

	if err != nil {
		return "", fmt.Errorf("write to cs failed: %s [%s]", cargo.LogName(), err)
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
