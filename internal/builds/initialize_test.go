package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/testdata"
)

func TestInitialize(t *testing.T) {
	suite.Run(t, new(SuiteInitialize))
}

type SuiteInitialize struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitialize) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitialize) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitialize) target() *Config {
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

func (this *SuiteInitialize) TestInitialize() {
	context, errs := Initialize(this.target())
	assert.Len(this.T(), errs, 0)
	assert.NotNil(this.T(), context)
	assert.NotEmpty(this.T(), context.Element)
	assert.NotEmpty(this.T(), context.Generate)
	assert.NotEmpty(this.T(), context.Encoding)
	assert.NotEmpty(this.T(), context.Poststep)
}
