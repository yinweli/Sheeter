package builds

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yinweli/Sheeter/internal/util"
)

// json-go讀取器模板代碼
const jsonGoReaderCode = `// generated by {{$.AppName}}, DO NOT EDIT.

package {{$.Namespace}}

import "encoding/json"

type {{$.ReaderName}} map[string]{{$.StructName}}

func (this *{{$.ReaderName}}) JsonPath() string {
	return "{{$.JsonPath}}"
}

func (this *{{$.ReaderName}}) FromJson(data []byte) error {
    return json.Unmarshal(data, this)
}
`

// writeJsonGo 輸出json-go代碼
func writeJsonGo(content *Content) error {
	if err := os.MkdirAll(filepath.Dir(content.JsonGoPath()), os.ModePerm); err != nil {
		return fmt.Errorf("%s: write json go failed: %w", content.ShowName(), err)
	} // if

	options := []string{
		"--src", content.SchemaPath(),
		"--src-lang", "json",
		"--out", content.JsonGoPath(),
		"--lang", "go",
		"--top-level", content.StructName(),
		"--package", content.Namespace(),
		"--just-types-and-package",
	}

	if err := util.ShellRun("quicktype", options...); err != nil {
		return fmt.Errorf("%s: write json go failed: %w", content.ShowName(), err)
	} // if

	if err := util.ShellRun("go", "fmt", content.JsonGoPath()); err != nil {
		return fmt.Errorf("%s: write json go failed: %w", content.ShowName(), err)
	} // if

	return nil
}

// writeJsonGoReader 輸出json-go讀取器, 由於quicktype對於結構命名有不一致的問題, 所以採取資料結構由quicktype執行, 而資料列表由模板執行的方式
func writeJsonGoReader(content *Content) error {
	if err := util.TmplWrite(content.JsonGoReaderPath(), jsonGoReaderCode, content, content.Bom); err != nil {
		return fmt.Errorf("%s: write json go reader failed: %w", content.ShowName(), err)
	} // if

	if err := util.ShellRun("go", "fmt", content.JsonGoReaderPath()); err != nil {
		return fmt.Errorf("%s: write json go reader failed: %w", content.ShowName(), err)
	} // if

	return nil
}
