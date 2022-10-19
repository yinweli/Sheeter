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

func TestGenerate(t *testing.T) {
	suite.Run(t, new(SuiteGenerate))
}

type SuiteGenerate struct {
	suite.Suite
	workDir string
}

func (this *SuiteGenerate) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteGenerate) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteGenerate) target() *Config {
	target := &Config{
		Global: Global{
			ExportJson:      true,
			ExportProto:     true,
			SimpleNamespace: false,
			LineOfName:      1,
			LineOfNote:      2,
			LineOfField:     3,
			LineOfLayer:     4,
			LineOfData:      5,
		},
		Elements: []Element{
			{Excel: testdata.ExcelReal, Sheet: testdata.SheetData},
		},
	}
	return target
}

func (this *SuiteGenerate) TestGenerate() {
	context, errs := Initialize(this.target())
	assert.Len(this.T(), errs, 0)
	assert.NotNil(this.T(), context)
	errs = Generate(context)
	assert.Len(this.T(), errs, 0)
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "realData.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "realDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.SchemaPath, "realData.proto"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.CsPath, "realDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.GoPath, "realDataReader.go"))
}
