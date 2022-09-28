package tmpls

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
)

// Initialize 初始化
func Initialize(cmd *cobra.Command) error {
	flags := cmd.Flags()

	if flags.Changed(flagClean) {
		if clean, err := flags.GetBool(flagClean); err == nil && clean {
			_ = os.RemoveAll(internal.PathTmpl)
		} // if
	} // if

	tmpls := []*Tmpl{
		JsonCsStruct,
		JsonCsReader,
		JsonGoStruct,
		JsonGoReader,
		ProtoSchema,
		ProtoCsReader,
		ProtoGoReader,
		ProtoCsBat,
		ProtoCsSh,
		ProtoGoBat,
		ProtoGoSh,
	}

	for _, itor := range tmpls {
		if err := itor.load(); err != nil {
			return fmt.Errorf("tmpl initialize failed: %w", err)
		} // if

		if err := itor.save(); err != nil {
			return fmt.Errorf("tmpl initialize failed: %w", err)
		} // if
	} // for

	return nil
}

// Tmpl 模板資料
type Tmpl struct {
	Name string // 模板檔名
	Data string // 模板字串
}

// load 讀取模板
func (this *Tmpl) load() error {
	path := this.path()

	if utils.ExistFile(path) == false {
		return nil
	} // if

	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("%s: tmpl load failed: %w", this.Name, err)
	} // if

	this.Data = string(data)
	return nil
}

// save 儲存模板
func (this *Tmpl) save() error {
	if err := utils.WriteFile(this.path(), []byte(this.Data)); err != nil {
		return fmt.Errorf("%s: tmpl save failed: %w", this.Name, err)
	} // if

	return nil
}

// path 取得模板路徑
func (this *Tmpl) path() string {
	return filepath.Join(internal.PathTmpl, this.Name)
}

// JsonCsStruct json-cs結構模板
var JsonCsStruct = &Tmpl{
	Name: internal.FileTmplJsonCsStruct,
	Data: `// generated by {{$.AppName}}, DO NOT EDIT.

using Newtonsoft.Json;
using System.Collections.Generic;

namespace {{$.NamespaceJson}} {
    public partial class {{$.StructName}} {
{{- range $.Fields}}
        // {{$.FieldNote .}}
        [JsonProperty("{{$.FieldName .}}")]
        public {{$.FieldTypeCs .}} {{$.FieldName .}} { get; set; }
{{- end}}
    }

    public partial class {{$.StorerName}} {
        public Dictionary<{{$.PkeyTypeCs}}, {{$.StructName}}> {{$.StorerDatas}} = new Dictionary<{{$.PkeyTypeCs}}, {{$.StructName}}>(); 
    }
}
`,
}

// JsonCsReader json-cs讀取器模板
var JsonCsReader = &Tmpl{
	Name: internal.FileTmplJsonCsReader,
	Data: `// generated by {{$.AppName}}, DO NOT EDIT.

using Newtonsoft.Json;
using System;
using System.IO;
using System.Collections.Generic;

namespace {{$.NamespaceJson}} {
    public partial class {{$.ReaderName}} {
        public static string FileName() {
            return "{{$.FileJsonData}}";
        }

        public bool FromPathFull(string path) {
            return FromData(File.ReadAllText(path));
        }

        public bool FromPathHalf(string path) {
            return FromData(File.ReadAllText(Path.Combine(path, FileName())));
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<{{$.StorerName}}>(data);
            return Datas != null;
        }

        public Dictionary<{{$.PkeyTypeCs}}, {{$.StructName}}> Data {
            get {
                return Datas.{{$.StorerDatas}};
            }
        }

        private {{$.StorerName}} Datas = null;
    }
}
`,
}

// JsonGoStruct json-go結構模板
var JsonGoStruct = &Tmpl{
	Name: internal.FileTmplJsonGoStruct,
	Data: `// generated by {{$.AppName}}, DO NOT EDIT.

package {{$.NamespaceJson}}

type {{$.StructName}} struct {
{{- range $.Fields}}
	// {{$.FieldNote .}}
	{{$.FieldName .}} {{$.FieldTypeGo .}} ` + "`json:\"{{$.FieldName .}}\"`" + `
{{- end}}
}

type {{$.StorerName}} struct {
	{{$.StorerDatas}} map[{{$.PkeyTypeGo}}]{{$.StructName}}
}
`,
}

// JsonGoReader json-go讀取器模板
var JsonGoReader = &Tmpl{
	Name: internal.FileTmplJsonGoReader,
	Data: `// generated by {{$.AppName}}, DO NOT EDIT.

package {{$.NamespaceJson}}

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type {{$.ReaderName}} struct {
	{{$.StorerName}}
}

func (this *{{$.ReaderName}}) FileName() string {
	return "{{$.FileJsonData}}"
}

func (this *{{$.ReaderName}}) FromPathFull(path string) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("{{$.ReaderName}}: from path full failed: %w", err)
	}

	return this.FromData(data)
}

func (this *{{$.ReaderName}}) FromPathHalf(path string) error {
	data, err := os.ReadFile(filepath.Join(path, this.FileName()))

	if err != nil {
		return fmt.Errorf("{{$.ReaderName}}: from path half failed: %w", err)
	}

	return this.FromData(data)
}

func (this *{{$.ReaderName}}) FromData(data []byte) error {
	this.{{$.StorerName}} = {{$.StorerName}}{
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]{{$.StructName}}{},
	}

	if err := json.Unmarshal(data, &this.{{$.StorerName}}); err != nil {
		return err
	}

	return nil
}
`,
}

// ProtoSchema proto架構模板
var ProtoSchema = &Tmpl{
	Name: internal.FileTmplProtoSchema,
	Data: `// generated by {{$.AppName}}, DO NOT EDIT.

syntax = "proto3";
package {{$.NamespaceProto}};
option go_package = ".;{{$.NamespaceProto}}";

{{- range $.Depend}}
import '{{$.ProtoDepend .}}';
{{- end}}

message {{$.StructName}} {
{{- range $i, $f := $.Fields}}
  {{$.FieldTypeProto .}} {{$.FieldName .}} = {{$.Add $i 1}}; // {{$.FieldNote .}}
{{- end}}
}

message {{$.StorerName}} {
  map<{{$.PkeyTypeProto}}, {{$.StructName}}> Datas = 1;
}
`,
}

// ProtoCsReader proto-cs讀取器模板
var ProtoCsReader = &Tmpl{
	Name: internal.FileTmplProtoCsReader,
	Data: "", // TODO: proto-cs讀取器模板
}

// ProtoGoReader proto-go讀取器模板
var ProtoGoReader = &Tmpl{
	Name: internal.FileTmplProtoGoReader,
	Data: "", // TODO: proto-go讀取器模板
}

// ProtoCsBat proto-cs-bat模板
var ProtoCsBat = &Tmpl{
	Name: internal.FileTmplProtoCsBat,
	Data: `REM generated by {{$.AppName}}, DO NOT EDIT.
REM please run this on PowerShell / Windows PowerShell 
REM https://learn.microsoft.com/zh-tw/powershell/scripting/whats-new/what-s-new-in-powershell-73?view=powershell-7.2

mkdir {{.PathProtoCs}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.PathProtoSchema}} --csharp_out=./{{.PathProtoCs}} ./{{.PathProtoName}}
{{- end}}
`,
}

// ProtoCsSh proto-cs-sh模板
var ProtoCsSh = &Tmpl{
	Name: internal.FileTmplProtoCsSh,
	Data: `# generated by {{$.AppName}}, DO NOT EDIT.

mkdir {{.PathProtoCs}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.PathProtoSchema}} --csharp_out=./{{.PathProtoCs}} ./{{.PathProtoName}}
{{- end}}
`,
}

// ProtoGoBat proto-go-bat模板
var ProtoGoBat = &Tmpl{
	Name: internal.FileTmplProtoGoBat,
	Data: `REM generated by {{$.AppName}}, DO NOT EDIT.
REM please run this on PowerShell / Windows PowerShell 
REM https://learn.microsoft.com/zh-tw/powershell/scripting/whats-new/what-s-new-in-powershell-73?view=powershell-7.2

mkdir {{.PathProtoGo}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.PathProtoSchema}} --go_out=./{{.PathProtoGo}} ./{{.PathProtoName}}
{{- end}}
`,
}

// ProtoGoSh proto-go-sh模板
var ProtoGoSh = &Tmpl{
	Name: internal.FileTmplProtoGoSh,
	Data: `# generated by {{$.AppName}}, DO NOT EDIT.

mkdir {{.PathProtoGo}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.PathProtoSchema}} --go_out=./{{.PathProtoGo}} ./{{.PathProtoName}}
{{- end}}
`,
}
