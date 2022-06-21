package core

import (
	"fmt"
	"os"
	"path"

	"github.com/yinweli/Sheeter/internal/util"
)

// executeJsonCs 輸出json/cs
func (this *Task) executeJsonCs() error {
	err := os.MkdirAll(path.Dir(this.jsonCsFilePath()), os.ModePerm)

	if err != nil {
		return fmt.Errorf("generate c# failed: %s\n%s", this.originalName(), err)
	} // if

	err, detail := util.ShellRun("quicktype", []string{
		"--src", this.jsonFilePath(),
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
