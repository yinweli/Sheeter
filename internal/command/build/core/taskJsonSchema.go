package core

import (
	"fmt"

	"github.com/yinweli/Sheeter/internal/util"
)

// runJsonSchema 輸出json架構
func (this *Task) runJsonSchema() error {
	obj := util.JsonObj{}

	for _, itor := range this.columns {
		if itor.Field.IsShow() {
			obj[itor.Name] = itor.Field.ToJsonDefault()
		} // if
	} // for

	err := util.JsonWrite(this.jsonSchemaFilePath(), obj, this.global.Bom)

	if err != nil {
		return fmt.Errorf("generate schema json failed: %s\n%w", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
