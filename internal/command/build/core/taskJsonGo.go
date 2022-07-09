package core

import (
	"fmt"
	"os"
	"path"

	"github.com/yinweli/Sheeter/internal/util"
)

// runJsonGo 輸出json-go
func (this *Task) runJsonGo() error {
	err := os.MkdirAll(path.Dir(this.jsonGoFilePath()), os.ModePerm)

	if err != nil {
		return fmt.Errorf("generate go failed: %s\n%s", this.originalName(), err)
	} // if

	err, detail := util.ShellRun("quicktype", []string{
		"--src", this.jsonSchemaFilePath(),
		"--src-lang", "json",
		"--out", this.jsonGoFilePath(),
		"--lang", "go",
		"--top-level", this.structName(),
		"--package", this.namespace(),
		"--just-types-and-package",
	}...)

	if err != nil {
		return fmt.Errorf("generate go failed: %s\n%s", this.originalName(), detail)
	} // if

	err, detail = util.ShellRun("go", "fmt", this.jsonGoFilePath())

	if err != nil {
		return fmt.Errorf("generate go failed: %s\n%s", this.originalName(), detail)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
