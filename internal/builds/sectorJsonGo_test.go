package builds

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestSectorJsonGo(t *testing.T) {
	suite.Run(t, new(SuiteSectorJsonGo))
}

type SuiteSectorJsonGo struct {
	suite.Suite
	workDir string
	code    []byte
	reader  []byte
}

func (this *SuiteSectorJsonGo) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.code = []byte(`package realdata

type Struct struct {
	S     S     ` + "`json:\"S\"`" + `
	Name0 int64 ` + "`json:\"name0\"`" + `
}

type S struct {
	A     []A  ` + "`json:\"A\"`" + `
	Name1 bool ` + "`json:\"name1\"`" + `
}

type A struct {
	Name2 int64  ` + "`json:\"name2\"`" + `
	Name3 string ` + "`json:\"name3\"`" + `
}
`)
	this.reader = []byte(fmt.Sprintf(`// generated by sheeter, DO NOT EDIT.

package realdata

import (
	"encoding/json"
	"os"
	"strconv"
)

type Reader map[int64]Struct

var Json = "%s"

func FromJsonFile(path string) (reader Reader, err error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return FromJsonBytes(data)
}

func FromJsonBytes(data []byte) (reader Reader, err error) {
	temps := map[string]Struct{}

	if err := json.Unmarshal(data, &temps); err != nil {
		return nil, err
	}

	datas := Reader{}

	for key, value := range temps {
		k, err := strconv.ParseInt(key, 10, 64)

		if err != nil {
			return nil, err
		}

		datas[k] = value
	}

	return datas, nil
}
`, filepath.ToSlash(filepath.Join("json", "realData.json"))))
}

func (this *SuiteSectorJsonGo) TearDownSuite() {
	_ = os.RemoveAll(internal.PathCode)
	_ = os.RemoveAll(internal.PathJsonSchema)
	_ = os.RemoveAll(internal.PathJsonGo)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteSectorJsonGo) target() *Sector {
	target := &Sector{
		Global: Global{
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
		},
		Element: Element{
			Excel: testdata.ExcelNameReal,
			Sheet: testdata.SheetName,
		},
	}
	return target
}

func (this *SuiteSectorJsonGo) TestSectorJsonGoCode() {
	target := this.target()
	assert.Nil(this.T(), SectorInit(target))
	assert.Nil(this.T(), SectorJsonSchema(target))
	assert.Nil(this.T(), SectorJsonGoCode(target))
	testdata.CompareFile(this.T(), target.FileJsonGoCode(), this.code)
	target.Close()
}

func (this *SuiteSectorJsonGo) TestSectorJsonGoReader() {
	target := this.target()
	assert.Nil(this.T(), SectorInit(target))
	assert.Nil(this.T(), SectorJsonSchema(target))
	assert.Nil(this.T(), SectorJsonGoReader(target))
	testdata.CompareFile(this.T(), target.FileJsonGoReader(), this.reader)
	target.Close()
}
