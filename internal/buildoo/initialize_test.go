package buildoo

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

func (this *SuiteInitialize) target() *Config {
	target := &Config{
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
				Excel: "test1",
				Sheet: "data1",
			},
			{
				Excel: "test2",
				Sheet: "data2",
			},
		},
	}
	return target
}

func (this *SuiteInitialize) TestInitialize() {
	target := this.target()
	context := Initialize(target)
	assert.NotNil(this.T(), context)
	assert.Equal(this.T(), &target.Global, context.Global)
	assert.Len(this.T(), context.element, len(target.Elements))
	element0, ok := context.element[0].(*initializeElement)
	assert.True(this.T(), ok)
	assert.NotNil(this.T(), element0)
	assert.Equal(this.T(), &target.Global, element0.Global)
	assert.Equal(this.T(), target.Elements[0], element0.Element)
	element1, ok := context.element[1].(*initializeElement)
	assert.True(this.T(), ok)
	assert.NotNil(this.T(), element1)
	assert.Equal(this.T(), &target.Global, element1.Global)
	assert.Equal(this.T(), target.Elements[1], element1.Element)
}
