package tasks

import (
	"fmt"
	"os"
	"path"

	"github.com/yinweli/Sheeter/internal/util"
)

// runJsonCs 輸出json-cs, 由於quicktype對於結構命名有不一致的問題, 所以採取資料結構由quicktype執行, 而資料列表由模板執行的方式
func (this *Task) runJsonCs() error {
	err := os.MkdirAll(path.Dir(this.jsonCsFilePath()), os.ModePerm)

	if err != nil {
		return fmt.Errorf("generate c# failed: %s\n%w", this.originalName(), err)
	} // if

	err = util.ShellRun("quicktype", []string{
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
		return fmt.Errorf("generate c# failed: %s\n%w", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
