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
	_ = os.RemoveAll(internal.EnumPath)
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteGenerate) target() *Config {
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

func (this *SuiteGenerate) TestGenerate() {
	context, errs := Initialize(this.target())
	assert.Empty(this.T(), errs)
	assert.Empty(this.T(), Generate(context))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "RealData.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.SchemaPath, "realData.proto"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(internal.EnumPath, internal.SchemaPath, "realEnum.proto"))
}
