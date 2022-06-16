package core

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// jsonCsCode json/c#程式碼模板
const jsonCsCode = `// generation by sheeter ^o<, from {{$.OriginalName}}
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Text;

namespace {{$.CsNamespace}} {
    public class {{$.StructName}} { {{$.SetLine}}
        public const string fileName = "{{$.JsonFileName}}";
{{range .Columns}}{{if .Field.IsShow}}        public {{.Field.TypeCs}} {{.Name}}; // {{.Note}}{{$.NewLine}}{{end}}{{end}}

        public static Dictionary<int, {{$.StructName}}> Parse(string s) {
            return JsonConvert.DeserializeObject<Dictionary<int, {{$.StructName}}>>(s);
        }

        public static Dictionary<int, {{$.StructName}}> Parse(byte[] b)
        {
            return Parse(Encoding.UTF8.GetString(b));
        }
    }
} // namespace {{$.CsNamespace}}
`

// executeJsonCs 輸出json/cs
func (this *Task) executeJsonCs() error {
	stemplateCode := STemplateCode{
		STemplate: STemplate{
			OriginalName: this.originalName(),
			StructName:   this.structName(),
		},
		JsonFileName: this.jsonFileName(),
		Columns:      this.columns,
	}
	bytes, err := stemplateCode.Generate(jsonCsCode)

	if err != nil {
		return fmt.Errorf("generate cs failed: %s [%s]", this.originalName(), err)
	} // if

	err = util.FileWrite(this.jsonCsFilePath(), bytes, this.global.Bom)

	if err != nil {
		return fmt.Errorf("write to cs failed: %s [%s]", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.IncrBy(taskProgressM)
	} // if

	return nil
}
