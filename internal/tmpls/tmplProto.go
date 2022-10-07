package tmpls

import (
	"github.com/yinweli/Sheeter/internal"
)

// ProtoSchema proto架構模板
var ProtoSchema = &Tmpl{
	Name: internal.TmplProtoSchemaFile,
	Data: HeaderCode + `
syntax = "proto3";
package {{$.ProtoNamespace $.SimpleNamespace}};
option go_package = ".;{{$.ProtoNamespace $.SimpleNamespace}}";

{{- range $.Depend}}
import '{{$.ProtoDepend .}}';
{{- end}}

message {{$.StructName}} {
{{- range $i, $f := $.Fields}}
  {{$.FieldTypeProto .}} {{$.FieldName .}} = {{$.Add $i 1}}; // {{$.FieldNote .}}
{{- end}}
}
{{- if $.Reader}}

message {{$.StorerName}} {
  map<{{$.PkeyTypeProto}}, {{$.StructName}}> Datas = 1;
}
{{- end}}
`,
}

// ProtoCsReader proto-cs讀取器模板
var ProtoCsReader = &Tmpl{
	Name: internal.TmplProtoCsReaderFile,
	Data: HeaderCode + `
using System.Collections.Generic;

namespace {{$.ProtoNamespace $.SimpleNamespace | $.FirstUpper}} {
    public partial class {{$.ReaderName}} {
        public string DataName() {
            return "{{$.ProtoDataName}}";
        }

        public string DataExt() {
            return "{{$.ProtoDataExt}}";
        }

        public string DataFile() {
            return "{{$.ProtoDataFile}}";
        }

        public bool FromData(byte[] data) {
            Datas = {{$.StorerName}}.Parser.ParseFrom(data);
            return Datas != null;
        }

        public {{$.PkeyTypeCs}}[] MergeData(byte[] data) {
            var repeats = new List<{{$.PkeyTypeCs}}>();
            var tmpl = {{$.StorerName}}.Parser.ParseFrom(data);

            if (tmpl == null)
                return repeats.ToArray();

            if (Datas == null)
                Datas = new {{$.StorerName}}();

            foreach (var itor in tmpl.{{$.StorerDatas}}) {
                if (Data.ContainsKey(itor.Key) == false)
                    Data[itor.Key] = itor.Value;
                else
                    repeats.Add(itor.Key);
            }

            return repeats.ToArray();
        }

        public IDictionary<{{$.PkeyTypeCs}}, {{$.StructName}}> Data {
            get {
                return Datas.{{$.StorerDatas}};
            }
        }

        private {{$.StorerName}} Datas = null;
    }
}
`,
}

// ProtoGoReader proto-go讀取器模板
var ProtoGoReader = &Tmpl{
	Name: internal.TmplProtoGoReaderFile,
	Data: HeaderCode + `
package {{$.ProtoNamespace $.SimpleNamespace}}

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type {{$.ReaderName}} struct {
	*{{$.StorerName}}
}

func (this *{{$.ReaderName}}) DataName() string {
	return "{{$.ProtoDataName}}"
}

func (this *{{$.ReaderName}}) DataExt() string {
	return "{{$.ProtoDataExt}}"
}

func (this *{{$.ReaderName}}) DataFile() string {
	return "{{$.ProtoDataFile}}"
}

func (this *{{$.ReaderName}}) FromData(data []byte) error {
	this.{{$.StorerName}} = &{{$.StorerName}}{
		Datas: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
	}

	if err := proto.Unmarshal(data, this.{{$.StorerName}}); err != nil {
		return fmt.Errorf("{{$.ReaderName}}: from data failed: %w", err)
	}

	return nil
}

func (this *{{$.ReaderName}}) MergeData(data []byte) (repeats []{{$.PkeyTypeGo}}) {
	tmpl := &{{$.StorerName}}{
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
		return repeats
	}

	if this.{{$.StorerName}} == nil {
		this.{{$.StorerName}} = &{{$.StorerName}}{
			{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
		}
	}

	for k, v := range tmpl.{{$.StorerDatas}} {
		if _, ok := this.{{$.StorerName}}.{{$.StorerDatas}}[k]; ok == false {
			this.{{$.StorerName}}.{{$.StorerDatas}}[k] = v
		} else {
			repeats = append(repeats, k)
		}
	}

	return repeats
}
`,
}

// ProtoCsBat proto-cs-bat模板
var ProtoCsBat = &Tmpl{
	Name: internal.TmplProtoCsBatFile,
	Data: HeaderBat + `
mkdir {{.ProtoCsPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --csharp_out=./{{.ProtoCsPath}} ./{{.ProtoPath}}
{{- end}}
`,
}

// ProtoCsSh proto-cs-sh模板
var ProtoCsSh = &Tmpl{
	Name: internal.TmplProtoCsShFile,
	Data: HeaderSh + `
mkdir {{.ProtoCsPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --csharp_out=./{{.ProtoCsPath}} ./{{.ProtoPath}}
{{- end}}
`,
}

// ProtoGoBat proto-go-bat模板
var ProtoGoBat = &Tmpl{
	Name: internal.TmplProtoGoBatFile,
	Data: HeaderBat + `
mkdir {{.ProtoGoPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --go_out=./{{.ProtoGoPath}} ./{{.ProtoPath}}
{{- end}}
`,
}

// ProtoGoSh proto-go-sh模板
var ProtoGoSh = &Tmpl{
	Name: internal.TmplProtoGoShFile,
	Data: HeaderSh + `
mkdir {{.ProtoGoPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --go_out=./{{.ProtoGoPath}} ./{{.ProtoPath}}
{{- end}}
`,
}
