package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/excels"
	"github.com/yinweli/Sheeter/testdata"
)

func TestGenerate(t *testing.T) {
	suite.Run(t, new(SuiteGenerate))
}

type SuiteGenerate struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteGenerate) SetupSuite() {
	this.Change("test-generate")
}

func (this *SuiteGenerate) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteGenerate) target() *Config {
	target := &Config{
		Global: Global{
			ExportJson:      true,
			ExportProto:     true,
			ExportEnum:      true,
			SimpleNamespace: false,
			LineOfTag:       1,
			LineOfName:      2,
			LineOfNote:      3,
			LineOfField:     4,
			LineOfLayer:     5,
			LineOfData:      6,
			LineOfEnum:      2,
			Tags:            "A",
		},
		Inputs: []string{
			testdata.ExcelReal + sheeter.SeparateSheet + testdata.SheetData,
			testdata.ExcelReal + sheeter.SeparateSheet + testdata.SheetEnum,
		},
	}
	return target
}

func (this *SuiteGenerate) TestGenerate() {
	context, errs := Initialize(this.target())
	assert.Empty(this.T(), errs)
	assert.Empty(this.T(), Generate(context))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "RealData.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "realData.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.SchemaPath, "realData.proto"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.CsPath, "RealDataReader.cs"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.GoPath, "realDataReader.go"))
	assert.FileExists(this.T(), filepath.Join(sheeter.EnumPath, sheeter.SchemaPath, "realEnum.proto"))
}
