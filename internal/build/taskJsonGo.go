package build

import (
	"fmt"
	"os"
	"path"

	"github.com/yinweli/Sheeter/internal/util"
)

// runJsonGo 輸出json-go, 由於quicktype對於結構命名有不一致的問題, 所以採取資料結構由quicktype執行, 而資料列表由模板執行的方式
func (this *Task) runJsonGo() error {
	err := os.MkdirAll(path.Dir(this.jsonGoFilePath()), os.ModePerm)

	if err != nil {
		return fmt.Errorf("generate go failed: %s\n%w", this.originalName(), err)
	} // if

	err = util.ShellRun("quicktype", []string{
		"--src", this.jsonSchemaFilePath(),
		"--src-lang", "json",
		"--out", this.jsonGoFilePath(),
		"--lang", "go",
		"--top-level", this.structName(),
		"--package", this.namespace(),
		"--just-types-and-package",
	}...)

	if err != nil {
		return fmt.Errorf("generate go failed: %s\n%w", this.originalName(), err)
	} // if

	err = util.ShellRun("go", "fmt", this.jsonGoFilePath())

	if err != nil {
		return fmt.Errorf("generate go failed: %s\n%w", this.originalName(), err)
	} // if

	if this.bar != nil {
		this.bar.Increment()
	} // if

	return nil
}
