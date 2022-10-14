package builds

import (
	"os"
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

func (this *SuiteGenerate) target() *Context {
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

func (this *SuiteGenerate) TestGenerate() {
	target := this.target()
	assert.Empty(this.T(), Initialize(target))
	assert.Empty(this.T(), Generate(target))
}
