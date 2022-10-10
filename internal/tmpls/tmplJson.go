package tmpls

import (
	"github.com/yinweli/Sheeter/internal"
)

// JsonCsStruct json-cs結構模板
var JsonCsStruct = &Tmpl{
	Name: internal.TmplJsonCsStructFile,
	Data: HeaderCode + `
using Newtonsoft.Json;
using System.Collections.Generic;

namespace {{$.JsonNamespace $.SimpleNamespace | $.FirstUpper}} {
    public partial class {{$.StructName}} {
{{- range $.Fields}}
        // {{$.FieldNote .}}
        [JsonProperty("{{$.FieldName .}}")]
        public {{$.FieldTypeCs .}} {{$.FieldName .}} { get; set; }
{{- end}}
    }
{{- if $.Reader}}

    public partial class {{$.StorerName}} {
        public Dictionary<{{$.PkeyTypeCs}}, {{$.StructName}}> {{$.StorerDatas}} = new Dictionary<{{$.PkeyTypeCs}}, {{$.StructName}}>(); 
    }
{{- end}}
}
`,
}

// JsonCsReader json-cs讀取器模板
var JsonCsReader = &Tmpl{
	Name: internal.TmplJsonCsReaderFile,
	Data: HeaderCode + `
using Newtonsoft.Json;
using System.Collections.Generic;

namespace {{$.JsonNamespace $.SimpleNamespace | $.FirstUpper}} {
    using Data_ = {{$.StructName}};
    using PKey_ = {{$.PkeyTypeCs}};
    using Storer_ = {{$.StorerName}};

    public partial class {{$.ReaderName}} : ReaderInterface {
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

        public bool ContainsKey(PKey_ key) {
            return storer.{{$.StorerDatas}}.ContainsKey(key);
        }

        public bool TryGetValue(PKey_ key, out Data_ value) {
            return storer.{{$.StorerDatas}}.TryGetValue(key, out value);
        }

        public IEnumerator<KeyValuePair<PKey_, Data_>> GetEnumerator() {
            return storer.{{$.StorerDatas}}.GetEnumerator();
        }

        public string DataName() {
            return "{{$.JsonDataName}}";
        }

        public string DataExt() {
            return "{{$.JsonDataExt}}";
        }

        public string DataFile() {
            return "{{$.JsonDataFile}}";
        }

        public string FromData(string data) {
            Storer_ result;

            try {
                result = JsonConvert.DeserializeObject<Storer_>(data);
            } catch {
                return "from data failed: deserialize failed";
            }

            if (result == null)
                return "from data failed: result null";

            storer = result;
            return string.Empty;
        }

        public string MergeData(string data) {
            Storer_ result;

            try {
                result = JsonConvert.DeserializeObject<Storer_>(data);
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

        private Storer_ storer = new Storer_();
    }
}
`,
}

// JsonCsDepot json-cs倉庫模板
var JsonCsDepot = &Tmpl{
	Name: internal.TmplJsonCsDepotFile, // TODO: 產生讀取器列表時, 必須把不會產生讀取器的Struct過濾掉
	Data: HeaderCode + `
using System.Collections.Generic;

namespace {{$.JsonNamespace $.SimpleNamespace | $.FirstUpper}} {
    public partial class Depot {
{{- range $.Struct}}
{{- if .Reader}}
        public readonly {{.ReaderName}} {{.StructName}} = new {{.ReaderName}}();
{{- end}}
{{- end}}
        private readonly List<ReaderInterface> Readers = new List<ReaderInterface>();
        
        public Depot() {
{{- range $.Struct}}
{{- if .Reader}}
            Readers.Add({{.StructName}});
{{- end}}
{{- end}}
        }

        public bool FromData(DelegateLoad load, DelegateError error) {
            var result = true;

            foreach (var itor in Readers) {
                var data = load(itor.DataName(), itor.DataExt());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.FromData(data);

                if (message.Length != 0) {
                    result = false;
                    error(itor.DataName(), message);
                }
            }

            return result;
        }

        public bool MergeData(DelegateLoad load, DelegateError error) {
            var result = true;

            foreach (var itor in Readers) {
                var data = load(itor.DataName(), itor.DataExt());

                if (data == null || data.Length == 0)
                    continue;

                var message = itor.MergeData(data);

                if (message.Length != 0) {
                    result = false;
                    error(itor.DataName(), message);
                }
            }

            return result;
        }

        public delegate void DelegateError(string name, string message);
        public delegate string DelegateLoad(string name, string ext);
    }

    public interface ReaderInterface {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(string data);
        public string MergeData(string data);
    }
}
`,
}

// JsonGoStruct json-go結構模板
var JsonGoStruct = &Tmpl{
	Name: internal.TmplJsonGoStructFile,
	Data: HeaderCode + `
package {{$.JsonNamespace $.SimpleNamespace}}

type {{$.StructName}} struct {
{{- range $.Fields}}
	// {{$.FieldNote .}}
	{{$.FieldName .}} {{$.FieldTypeGo .}} ` + "`json:\"{{$.FieldName .}}\"`" + `
{{- end}}
}
{{- if $.Reader}}

type {{$.StorerName}} struct {
	{{$.StorerDatas}} map[{{$.PkeyTypeGo}}]{{$.StructName}}
}
{{- end}}
`,
}

// JsonGoReader json-go讀取器模板
var JsonGoReader = &Tmpl{
	Name: internal.TmplJsonGoReaderFile,
	Data: HeaderCode + `
package {{$.JsonNamespace $.SimpleNamespace}}

import (
	"encoding/json"
	"fmt"
)

type {{$.ReaderName}} struct {
	*{{$.StorerName}}
}

func (this *{{$.ReaderName}}) DataName() string {
	return "{{$.JsonDataName}}"
}

func (this *{{$.ReaderName}}) DataExt() string {
	return "{{$.JsonDataExt}}"
}

func (this *{{$.ReaderName}}) DataFile() string {
	return "{{$.JsonDataFile}}"
}

func (this *{{$.ReaderName}}) FromData(data []byte) error {
	this.{{$.StorerName}} = &{{$.StorerName}}{
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]{{$.StructName}}{},
	}

	if err := json.Unmarshal(data, this.{{$.StorerName}}); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *{{$.ReaderName}}) MergeData(data []byte) error {
	tmpl := &{{$.StorerName}}{
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]{{$.StructName}}{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return fmt.Errorf("merge data failed: %w", err)
	}

	if this.{{$.StorerName}} == nil {
		this.{{$.StorerName}} = &{{$.StorerName}}{
			{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]{{$.StructName}}{},
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
`,
}

// JsonGoDepot json-go倉庫模板 // TODO: json-go倉庫模板
var JsonGoDepot = &Tmpl{
	Name: internal.TmplJsonGoDepotFile,
	Data: HeaderCode + `
package {{$.JsonNamespace $.SimpleNamespace}}

type Depot struct {
{{- range $.Struct}}
{{- if .Reader}}
	{{.StructName}} {{.ReaderName}}
{{- end}}
{{- end}}
	readers []ReaderInterface
}

func (this *Depot) FromData(load DepotLoad, error DepotError) bool {
	this.build()
	result := true

	for _, itor := range this.readers {
		data := load(itor.DataName(), itor.DataExt())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) MergeData(load DepotLoad, error DepotError) bool {
	this.build()
	result := true

	for _, itor := range this.readers {
		data := load(itor.DataName(), itor.DataExt())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) build() {
	if len(this.readers) == 0 {
		this.readers = append(
			this.readers,
{{- range $.Struct}}
{{- if .Reader}}
			&this.{{.StructName}},
{{- end}}
{{- end}}
		)
	}
}

type DepotError func(name string, err error)
type DepotLoad func(name, ext string) []byte

type ReaderInterface interface {
	DataName() string
	DataExt() string
	DataFile() string
	FromData(data []byte) error
	MergeData(data []byte) error
}
`,
}
