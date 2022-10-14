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
				LineOfField: 1,
				LineOfLayer: 2,
				LineOfNote:  3,
				LineOfData:  4,
			},
		},
		Sector: []*ContextSector{
			{
				Element: Element{
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
