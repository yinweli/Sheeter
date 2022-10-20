package tmpls

import (
	"github.com/yinweli/Sheeter/internal"
)

// JsonStructCs json結構cs模板
var JsonStructCs = &Tmpl{
	Name: internal.TmplJsonStructCsFile,
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

// JsonReaderCs json讀取器cs模板
var JsonReaderCs = &Tmpl{
	Name: internal.TmplJsonReaderCsFile,
	Data: HeaderCode + `
using Newtonsoft.Json;
using System.Collections.Generic;

namespace {{$.JsonNamespace $.SimpleNamespace | $.FirstUpper}} {
    using Data_ = {{$.StructName}};
    using PKey_ = {{$.PkeyTypeCs}};
    using Storer_ = {{$.StorerName}};

    public partial class {{$.ReaderName}} : Reader {
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

// JsonDepotCs json倉庫cs模板
var JsonDepotCs = &Tmpl{
	Name: internal.TmplJsonDepotCsFile,
	Data: HeaderCode + `
using System.Collections.Generic;

namespace {{$.JsonNamespace $.SimpleNamespace | $.FirstUpper}} {
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
        public string Load(string name, string ext, string fullname);
    }

    public interface Reader {
        public string DataName();
        public string DataExt();
        public string DataFile();
        public string FromData(string data);
        public string MergeData(string data);
        public void Clear();
    }
}
`,
}

// JsonStructGo json-go結構模板
var JsonStructGo = &Tmpl{
	Name: internal.TmplJsonStructGoFile,
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
	{{$.StorerDatas}} map[{{$.PkeyTypeGo}}]*{{$.StructName}}
}
{{- end}}
`,
}

// JsonReaderGo json-go讀取器模板
var JsonReaderGo = &Tmpl{
	Name: internal.TmplJsonReaderGoFile,
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
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
	}

	if err := json.Unmarshal(data, this.{{$.StorerName}}); err != nil {
		return fmt.Errorf("from data failed: %w", err)
	}

	return nil
}

func (this *{{$.ReaderName}}) MergeData(data []byte) error {
	tmpl := &{{$.StorerName}}{
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]*{{$.StructName}}{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
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

// JsonDepotGo json-go倉庫模板
var JsonDepotGo = &Tmpl{
	Name: internal.TmplJsonDepotGoFile,
	Data: HeaderCode + `
package {{$.JsonNamespace $.SimpleNamespace}}

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
