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
    public partial class {{$.ReaderName}} {
        public string DataName() {
            return "{{$.JsonDataName}}";
        }

        public string DataExt() {
            return "{{$.JsonDataExt}}";
        }

        public string DataFile() {
            return "{{$.JsonDataFile}}";
        }

        public bool FromData(string data) {
            Datas = JsonConvert.DeserializeObject<{{$.StorerName}}>(data);
            return Datas != null;
        }

        public {{$.PkeyTypeCs}}[] MergeData(string data) {
            var repeats = new List<{{$.PkeyTypeCs}}>();
            var tmpl = JsonConvert.DeserializeObject<{{$.StorerName}}>(data);

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

// JsonCsDepot json-cs倉庫模板 // TODO: json-cs倉庫模板
var JsonCsDepot = &Tmpl{
	Name: internal.TmplJsonCsDepotFile,
	Data: HeaderCode + `
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
		return fmt.Errorf("{{$.ReaderName}}: from data failed: %w", err)
	}

	return nil
}

func (this *{{$.ReaderName}}) MergeData(data []byte) (repeats []{{$.PkeyTypeGo}}) {
	tmpl := &{{$.StorerName}}{
		{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]{{$.StructName}}{},
	}

	if err := json.Unmarshal(data, tmpl); err != nil {
		return repeats
	}

	if this.{{$.StorerName}} == nil {
		this.{{$.StorerName}} = &{{$.StorerName}}{
			{{$.StorerDatas}}: map[{{$.PkeyTypeGo}}]{{$.StructName}}{},
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

// JsonGoDepot json-go倉庫模板 // TODO: json-go倉庫模板
var JsonGoDepot = &Tmpl{
	Name: internal.TmplJsonGoDepotFile,
	Data: HeaderCode + `
`,
}
