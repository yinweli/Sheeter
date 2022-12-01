package tmpls

import "github.com/yinweli/Sheeter/sheeter"

// ProtoSchema proto架構模板
var ProtoSchema = &Tmpl{
	Name: sheeter.TmplProtoSchemaFile,
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
	Name: sheeter.TmplProtoReaderCsFile,
	Data: HeaderCode + `
using System.Collections.Generic;

namespace {{$.ProtoNamespace $.SimpleNamespace | $.FirstUpper}} {
    using Data_ = {{$.StructName}};
    using PKey_ = {{$.PkeyTypeCs}};
    using Storer_ = {{$.StorerName}};

    public partial class {{$.ReaderName}} : Reader {
        public FileName FileName() {
            return new FileName("{{$.ProtoDataName}}", "{{$.ProtoDataExt}}");
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
	Name: sheeter.TmplProtoDepotCsFile,
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
                var filename = itor.FileName();
                var data = Loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(filename.File, message);
                }
            }

            return result;
        }

        public bool MergeData() {
            if (Loader == null)
                return false;

            var result = true;

            foreach (var itor in Readers) {
                var filename = itor.FileName();
                var data = Loader.Load(filename);

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    Loader.Error(filename.File, message);
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

    public class FileName {
        public FileName(string name, string ext) {
            this.name = name;
            this.ext = ext;
        }

        public string Name {
            get {
                return name;
            }
        }

        public string Ext {
            get {
                return ext;
            }
        }

        public string File {
            get {
                return name + ext;
            }
        }

        private readonly string name;
        private readonly string ext;
    }

    public interface Loader {
        public void Error(string name, string message);
        public byte[] Load(FileName filename);
    }

    public interface Reader {
        public FileName FileName();
        public string FromData(byte[] data);
        public string MergeData(byte[] data);
        public void Clear();
    }
}
`,
}

// ProtoReaderGo proto-go讀取器模板
var ProtoReaderGo = &Tmpl{
	Name: sheeter.TmplProtoReaderGoFile,
	Data: HeaderCode + `
package {{$.ProtoNamespace $.SimpleNamespace}}

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type {{$.ReaderName}} struct {
	*{{$.StorerName}}
}

func (this *{{$.ReaderName}}) FileName() FileName {
	return NewFileName("{{$.ProtoDataName}}", "{{$.ProtoDataExt}}")
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
	Name: sheeter.TmplProtoDepotGoFile,
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
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
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
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
		}
	}

	return result
}

func (this *Depot) Clear() {
	for _, itor := range this.readers {
		itor.Clear()
	}
}

type FileName struct {
	name string
	ext  string
}

func NewFileName(name, ext string) FileName {
	return FileName{
		name: name,
		ext:  ext,
	}
}

func (this FileName) Name() string {
	return this.name
}

func (this FileName) Ext() string {
	return this.ext
}

func (this FileName) File() string {
	return this.name + this.ext
}

type Loader interface {
	Error(name string, err error)
	Load(filename FileName) []byte
}

type Reader interface {
	FileName() FileName
	FromData(data []byte) error
	MergeData(data []byte) error
	Clear()
}
`,
}
