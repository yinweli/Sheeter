package builds

import (
	"os"
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

func (this *SuiteEncoding) target() *Context {
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

func (this *SuiteEncoding) TestEncoding() {
	target := this.target()
	assert.Empty(this.T(), Initialize(target))
	assert.Empty(this.T(), Generate(target))
	assert.Empty(this.T(), Encoding(target))
}
