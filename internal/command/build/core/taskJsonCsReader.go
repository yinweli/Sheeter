package core

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

const csReaderContent = `namespace {{$.Namespace}} {
    using System;
    using System.Collections.Generic;

    using Newtonsoft.Json;

    public partial class {{$.ReaderName}} {
        public static Dictionary<string, {{$.StructName}}> FromJson(string json) {
            return JsonConvert.DeserializeObject<Dictionary<string, {{$.StructName}}>>(json);
        }
    }
}`

// TmplCsReader json-cs讀取器模板資料
type TmplCsReader struct {
	Namespace  string // 命名空間
	StructName string // 結構名稱
	ReaderName string // 讀取器名稱
}

// runJsonCs 輸出json-cs讀取器
func (this *Task) runJsonCsReader() error {
	err := util.TmplWrite(this.jsonCsReaderFilePath(), this.global.Bom, csReaderContent, &TmplCsReader{
		Namespace:  this.namespace(),
		StructName: this.structName(),
		ReaderName: this.readerName(),
	})

	if err != nil {
		return fmt.Errorf("generate csReader failed: %s\n%s", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
