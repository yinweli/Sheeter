package tmpls

import "github.com/yinweli/Sheeter/sheeter"

// EnumSchema enum架構模板
var EnumSchema = &Tmpl{
	Name: sheeter.TmplEnumSchemaFile,
	Data: HeaderCode + `
syntax = "proto3";
package {{$.EnumNamespace $.SimpleNamespace}};
option go_package = ".;{{$.EnumNamespace $.SimpleNamespace}}";

enum {{$.StructName}} {
{{- range $.Enums}}
  {{$.FirstUpper .Name}} = {{.Index}}; // {{.Comment}}
{{- end}}
}
`,
}
