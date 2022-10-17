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
	_ = os.Remove(mixed.ProtoBatCsFile())
	_ = os.Remove(mixed.ProtoShCsFile())
	_ = os.Remove(mixed.ProtoBatGoFile())
	_ = os.Remove(mixed.ProtoShGoFile())
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
				LineOfName:  1,
				LineOfNote:  2,
				LineOfField: 3,
				LineOfLayer: 4,
				LineOfData:  5,
			},
			Elements: []Element{
				{
					Excel: testdata.ExcelNameReal,
					Sheet: testdata.SheetData,
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
	assert.Empty(this.T(), Poststep(target))
}
