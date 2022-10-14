package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/mixeds"
	"github.com/yinweli/Sheeter/testdata"
)

func TestPoststep(t *testing.T) {
	suite.Run(t, new(SuitePoststep))
}

type SuitePoststep struct {
	suite.Suite
	workDir string
}

func (this *SuitePoststep) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuitePoststep) TearDownSuite() {
	mixed := mixeds.NewMixed("", "")
	_ = os.Remove(mixed.ProtoCsBatFile())
	_ = os.Remove(mixed.ProtoCsShFile())
	_ = os.Remove(mixed.ProtoGoBatFile())
	_ = os.Remove(mixed.ProtoGoShFile())
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuitePoststep) target() *Context {
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

func (this *SuitePoststep) TestEncoding() {
	target := this.target()
	assert.Empty(this.T(), Initialize(target))
	assert.Empty(this.T(), Generate(target))
	assert.Empty(this.T(), Encoding(target))
	assert.Nil(this.T(), Poststep(target))
}
