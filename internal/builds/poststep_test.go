package builds

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststep(t *testing.T) {
	suite.Run(t, new(SuitePoststep))
}

type SuitePoststep struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuitePoststep) SetupSuite() {
	this.Change("test-poststep")
}

func (this *SuitePoststep) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuitePoststep) target() *Config {
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

func (this *SuitePoststep) TestPoststep() {
	context, errs := Initialize(this.target())
	assert.Empty(this.T(), errs)
	assert.Empty(this.T(), Generate(context))
	assert.Empty(this.T(), Encoding(context))
	assert.Empty(this.T(), Poststep(context))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.JsonPath, internal.GoPath, "depot.go"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.CsPath, "Depot.cs"))
	assert.FileExists(this.T(), filepath.Join(internal.ProtoPath, internal.GoPath, "depot.go"))
}
