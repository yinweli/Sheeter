package core

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// executeJsonSchema 輸出json架構
func (this *Task) executeJsonSchema() error {
	obj := make(util.JsonObj)

	for _, itor := range this.columns {
		if itor.Field.IsShow() {
			obj[itor.Name] = itor.Field.ToJsonDefault()
		} // if
	} // for

	err := util.JsonWrite(obj, this.jsonSchemaFilePath(), this.global.Bom)

	if err != nil {
		return fmt.Errorf("generate schema json failed: %s\n%s", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
