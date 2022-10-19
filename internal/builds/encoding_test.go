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
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncoding) target() *Config {
	target := &Config{
		Global: Global{
			ExportJson:      true,
			ExportProto:     true,
			Format:          true,
			SimpleNamespace: false,
			LineOfName:      1,
			LineOfNote:      2,
			LineOfField:     3,
			LineOfLayer:     4,
			LineOfData:      5,
		},
		Elements: []Element{
			{Excel: testdata.ExcelNameReal, Sheet: testdata.SheetData},
		},
	}
	return target
}

func (this *SuiteEncoding) TestEncoding() {
	context, errs := Initialize(this.target())
	assert.Len(this.T(), errs, 0)
	assert.NotNil(this.T(), context)
	errs = Generate(context)
	assert.Len(this.T(), errs, 0)
	errs = Encoding(context)
	assert.Len(this.T(), errs, 0)
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.DataPath, "realData.json"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.DataPath, "realData.bytes"))
}
