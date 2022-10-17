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

// ProtoReaderCs proto讀取器cs模板
var ProtoReaderCs = &Tmpl{
	Name: internal.TmplProtoReaderCsFile,
	Data: HeaderCode + `
using System.Collections.Generic;

namespace {{$.ProtoNamespace $.SimpleNamespace | $.FirstUpper}} {
    using Data_ = {{$.StructName}};
    using PKey_ = {{$.PkeyTypeCs}};
    using Storer_ = {{$.StorerName}};

    public partial class {{$.ReaderName}} : Reader {
        public string DataName() {
            return "{{$.ProtoDataName}}";
        }

        public string DataExt() {
            return "{{$.ProtoDataExt}}";
        }

        public string DataFile() {
            return "{{$.ProtoDataFile}}";
        }

        public string FromData(byte[] data) {
            Storer_ result;

            try {
                result = Storer_.Parser.ParseFrom(data);
            } catch {
                return "from data failed: deserialize failed";
            }

            if (result == null)
                return "from data failed: result null";

            storer = result;
            return string.Empty;
        }

        public string MergeData(byte[] data) {
            Storer_ result;

            try {
                result = Storer_.Parser.ParseFrom(data);
            } catch {
                return "merge data failed: deserialize failed";
            }

            if (result == null)
                return "merge data failed: result null";

            foreach (var itor in result.{{$.StorerDatas}}) {
                if (storer.{{$.StorerDatas}}.ContainsKey(itor.Key))
                    return "merge data failed: key repeat";

                storer.{{$.StorerDatas}}[itor.Key] = itor.Value;
            }

            return string.Empty;
        }

        public void Clear() {
            storer.{{$.StorerDatas}}.Clear();
        }

        public bool TryGetValue(PKey_ key, out Data_ value) {
            return storer.{{$.StorerDatas}}.TryGetValue(key, out value);
        }

        public bool ContainsKey(PKey_ key) {
            return storer.{{$.StorerDatas}}.ContainsKey(key);
        }

        public IEnumerator<KeyValuePair<PKey_, Data_>> GetEnumerator() {
            return storer.{{$.StorerDatas}}.GetEnumerator();
        }

        public Data_ this[PKey_ key] {
            get {
                return storer.{{$.StorerDatas}}[key];
            }
        }

        public ICollection<PKey_> Keys {
            get {
                return storer.{{$.StorerDatas}}.Keys;
            }
        }

        public ICollection<Data_> Values {
            get {
                return storer.{{$.StorerDatas}}.Values;
            }
        }

        public int Count {
            get {
                return storer.{{$.StorerDatas}}.Count;
            }
        }

        private Storer_ storer = new Storer_();
    }
}
`,
}

// ProtoDepotCs proto倉庫cs模板
var ProtoDepotCs = &Tmpl{
	Name: internal.TmplProtoDepotCsFile,
	Data: HeaderCode + `
using System.Collections.Generic;

namespace {{$.ProtoNamespace $.SimpleNamespace | $.FirstUpper}} {
    public partial class Depot {
        public Loader Loader { get; set; }
{{- range $.Struct}}
{{- if .Reader}}
        public readonly {{.ReaderName}} {{.StructName}} = new {{.ReaderName}}();
{{- end}}
{{- end}}
        private readonly List<Reader> Readers = new List<Reader>();

        public Depot() {
{{- range $.Struct}}
{{- if .Reader}}
            Readers.Add({{.StructName}});
{{- end}}
{{- end}}
        }

        public bool FromData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }

        public bool MergeData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var data = Loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(itor.DataName(), message);
                }
            }

            return result;
        }

        public void Clear() {
            foreach (var itor in Readers) {
                itor.Clear();
            }
        }
    }

    public interface Loader {
        public void Error(string name, string message);
        public byte[] Load(string name, string ext, string fullname);
    }

    public interface Reader {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(byte[] data);
        public string MergeData(byte[] data);
        public void Clear();
    }
}
`,
}

// ProtoReaderGo proto-go讀取器模板
var ProtoReaderGo = &Tmpl{
	Name: internal.TmplProtoReaderGoFile,
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
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
	}

	if err := proto.Unmarshal(data, this.{{$.StorerName}}); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *{{$.ReaderName}}) MergeData(data []byte) error {
	tmpl := &{{$.StorerName}}{
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
	}

	if err := proto.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.{{$.StorerName}} == nil {
		this.{{$.StorerName}} = &{{$.StorerName}}{
			{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
		}
	}

	for k, v := range tmpl.{{$.StorerDatas}} {
		if _, ok := this.{{$.StorerName}}.{{$.StorerDatas}}[k]; ok {
			return fmt.Errorf("merge data failed: key repeat")
		}

		this.{{$.StorerName}}.{{$.StorerDatas}}[k] = v
	}

	return nil
}

func (this *{{$.ReaderName}}) Clear() {
	this.{{$.StorerName}} = nil
}

func (this *{{$.ReaderName}}) Get(key {{$.PkeyTypeGo}}) (result *{{$.StructName}}, ok bool) {
	result, ok = this.{{$.StorerDatas}}[key]
	return result, ok
}

func (this *{{$.ReaderName}}) Keys() (result []{{$.PkeyTypeGo}}) {
	for itor := range this.{{$.StorerDatas}} {
		result = append(result, itor)
	}

	return result
}

func (this *{{$.ReaderName}}) Values() (result []*{{$.StructName}}) {
	for _, itor := range this.{{$.StorerDatas}} {
		result = append(result, itor)
	}

	return result
}

func (this *{{$.ReaderName}}) Count() int {
	return len(this.{{$.StorerDatas}})
}
`,
}

// ProtoDepotGo proto-go倉庫模板
var ProtoDepotGo = &Tmpl{
	Name: internal.TmplProtoDepotGoFile,
	Data: HeaderCode + `
package {{$.ProtoNamespace $.SimpleNamespace}}

func NewDepot(loader Loader) *Depot {
	depot := &Depot{}
	depot.loader = loader
	depot.readers = append(
		depot.readers,
{{- range $.Struct}}
{{- if .Reader}}
		&depot.{{.StructName}},
{{- end}}
{{- end}}
	)
	return depot
}

type Depot struct {
{{- range $.Struct}}
{{- if .Reader}}
	{{.StructName}} {{.ReaderName}}
{{- end}}
{{- end}}
	loader Loader
	readers []Reader
}

func (this *Depot) FromData() bool {
	if this.loader == nil {
		return false
	}

	result := true

	for _, itor := range this.readers {
		data := this.loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			this.loader.Error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) MergeData() bool {
	if this.loader == nil {
		return false
	}

	result := true

	for _, itor := range this.readers {
		data := this.loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			this.loader.Error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) Clear() {
	for _, itor := range this.readers {
		itor.Clear()
	}
}

type Loader interface {
	Error(name string, err error)
	Load(name, ext, fullname string) []byte
}

type Reader interface {
	DataName() string
	DataExt() string
	DataFile() string
	FromData(data []byte) error
	MergeData(data []byte) error
	Clear()
}
`,
}

// ProtoBatCs proto-bat-cs模板
var ProtoBatCs = &Tmpl{
	Name: internal.TmplProtoBatCsFile,
	Data: HeaderBat + `
mkdir {{.ProtoCsPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --csharp_out=./{{.ProtoCsPath}} ./{{.ProtoPath}}
{{- end}}
`,
}

// ProtoShCs proto-sh-cs模板
var ProtoShCs = &Tmpl{
	Name: internal.TmplProtoShCsFile,
	Data: HeaderSh + `
mkdir {{.ProtoCsPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --csharp_out=./{{.ProtoCsPath}} ./{{.ProtoPath}}
{{- end}}
`,
}

// ProtoBatGo proto-go-bat模板
var ProtoBatGo = &Tmpl{
	Name: internal.TmplProtoBatGoFile,
	Data: HeaderBat + `
mkdir {{.ProtoGoPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --go_out=./{{.ProtoGoPath}} ./{{.ProtoPath}}
{{- end}}
`,
}

// ProtoShGo proto-go-sh模板
var ProtoShGo = &Tmpl{
	Name: internal.TmplProtoShGoFile,
	Data: HeaderSh + `
mkdir {{.ProtoGoPath}}
{{- range $.Struct}}
protoc --experimental_allow_proto3_optional --proto_path=./{{.ProtoSchemaPath}} --go_out=./{{.ProtoGoPath}} ./{{.ProtoPath}}
{{- end}}
`,
}
