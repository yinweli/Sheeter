package builds

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal"
	"github.com/yinweli/Sheeter/internal/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestEncodingProto(t *testing.T) {
	suite.Run(t, new(SuiteEncodingProto))
}

type SuiteEncodingProto struct {
	suite.Suite
	workDir string
}

func (this *SuiteEncodingProto) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteEncodingProto) TearDownSuite() {
	_ = os.RemoveAll(internal.JsonPath)
	_ = os.RemoveAll(internal.ProtoPath)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteEncodingProto) target() *Config {
	target := &Config{
		Global: Global{
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
			LineOfData:  4,
		},
		Elements: []Element{
			{
				Excel: testdata.ExcelNameReal,
				Sheet: testdata.SheetName,
			},
		},
	}
	return target
}

func (this *SuiteEncodingProto) TestEncodingProto() {
	target := this.target()
	runtime := &Runtime{}
	assert.Nil(this.T(), Initialize(runtime, target))
	assert.Nil(this.T(), Generate(runtime, &Config{
		Global: Global{
			ExportProto:     true,
			SimpleNamespace: false,
		},
	}))
	assert.Nil(this.T(), encodingProto(runtime.Sector[0]))
	assert.True(this.T(), utils.ExistFile(runtime.Sector[0].ProtoDataPath()))
	runtime.Sector[0].Close()
}
