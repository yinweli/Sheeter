package core

import (
	"fmt"
	"os"
	"path"

	"github.com/yinweli/Sheeter/internal/util"
)

// runJsonCs 輸出json/cs
func (this *Task) runJsonCs() error {
	err := os.MkdirAll(path.Dir(this.jsonCsFilePath()), os.ModePerm)

	if err != nil {
		return fmt.Errorf("generate c# failed: %s\n%s", this.originalName(), err)
	} // if

	err, detail := util.ShellRun("quicktype", []string{
		"--src", this.jsonSchemaFilePath(),
		"--src-lang", "json",
		"--out", this.jsonCsFilePath(),
		"--lang", "cs",
		"--top-level", this.structName(),
		"--namespace", this.namespace(),
		"--array-type", "array",
		"--features", "attributes-only",
	}...)

	if err != nil {
		return fmt.Errorf("generate c# failed: %s\n%s", this.originalName(), detail)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
