package core

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// json-cs讀取器模板內容
const csReaderContent = `// generated by {{$.Namespace}}, DO NOT EDIT.

namespace {{$.Namespace}} {
    using System;
    using System.Collections.Generic;

    using Newtonsoft.Json;

    public partial class {{$.ReaderName}} {
		public static readonly string JsonFileName = "{{$.JsonFileName}}";

        public static Dictionary<string, {{$.StructName}}> FromJson(string data) {
            return JsonConvert.DeserializeObject<Dictionary<string, {{$.StructName}}>>(data);
        }
    }
}
`

// runJsonCsReader 輸出json-cs讀取器, 由於quicktype對於結構命名有不一致的問題, 所以採取資料結構由quicktype執行, 而資料列表由模板執行的方式
func (this *Task) runJsonCsReader() error {
	err := util.TmplWrite(this.jsonCsReaderFilePath(), csReaderContent, &TmplCsReader{
		JsonFileName: this.jsonFileName(),
		Namespace:    this.namespace(),
		StructName:   this.structName(),
		ReaderName:   this.readerName(),
	}, this.global.Bom)

	if err != nil {
		return fmt.Errorf("generate csReader failed: %s\n%w", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}

// TmplCsReader json-cs讀取器模板資料
type TmplCsReader struct {
	JsonFileName string // json檔名
	Namespace    string // 命名空間
	StructName   string // 結構名稱
	ReaderName   string // 讀取器名稱
}
