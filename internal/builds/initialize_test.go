package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitialize) target() *Context {
	target := &Context{
		Config: &Config{
			Global: Global{
				ExportJson:  true,
				ExportProto: true,
				LineOfName:  1,
				LineOfNote:  2,
				LineOfField: 3,
				LineOfLayer: 4,
				LineOfData:  5,
			},
			Elements: []Element{
				{
					Excel: testdata.ExcelNameReal,
					Sheet: testdata.SheetName,
				},
			},
		},
	}
	return target
}

func (this *SuiteInitialize) TestInitialize() {
	target := this.target()
	assert.Empty(this.T(), Initialize(target))
}
