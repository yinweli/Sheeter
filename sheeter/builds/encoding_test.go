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

func TestEncoding(t *testing.T) {
	suite.Run(t, new(SuiteEncoding))
}

type SuiteEncoding struct {
	suite.Suite
	testdata.TestEnv
}

func (this *SuiteEncoding) SetupSuite() {
	this.Change("test-encoding")
}

func (this *SuiteEncoding) TearDownSuite() {
	excels.CloseAll()
	this.Restore()
}

func (this *SuiteEncoding) target() *Config {
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

func (this *SuiteEncoding) TestEncoding() {
	context, errs := Initialize(this.target())
	assert.Empty(this.T(), errs)
	assert.Empty(this.T(), Generate(context))
	assert.Empty(this.T(), Encoding(context))
	assert.FileExists(this.T(), filepath.Join(sheeter.JsonPath, sheeter.DataPath, "realData.json"))
	assert.FileExists(this.T(), filepath.Join(sheeter.ProtoPath, sheeter.DataPath, "realData.bytes"))
}
