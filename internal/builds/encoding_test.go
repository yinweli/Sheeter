package builds

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncoding(t *testing.T) {
	suite.Run(t, new(SuiteEncoding))
}

type SuiteEncoding struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncoding) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncoding) TearDownSuite() {
	_ = os.RemoveAll(internal.EnumPath)
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncoding) target() *Config {
	target := &Config{
		Global: Global{
			ExportJson:      true,
			ExportProto:     true,
			ExportEnum:      true,
			SimpleNamespace: false,
			LineOfName:      1,
			LineOfNote:      2,
			LineOfField:     3,
			LineOfLayer:     4,
			LineOfData:      5,
			LineOfEnum:      2,
		},
		Inputs: []string{
			testdata.ExcelReal + internal.SeparateSheet + testdata.SheetData,
			testdata.ExcelReal + internal.SeparateSheet + testdata.SheetEnum,
		},
	}
	return target
}

func (this *SuiteEncoding) TestEncoding() {
	context, errs := Initialize(this.target())
	assert.Empty(this.T(), errs)
	assert.Empty(this.T(), Generate(context))
	assert.Empty(this.T(), Encoding(context))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.DataPath, "realData.json"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.DataPath, "realData.bytes"))
}
