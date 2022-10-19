package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeContext(t *testing.T) {
	suite.Run(t, new(SuiteInitializeContext))
}

type SuiteInitializeContext struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeContext) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeContext) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeContext) target() *Config {
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
			{Excel: "test1", Sheet: "data1"},
			{Excel: "test2", Sheet: "data2"},
		},
	}
	return target
}

func (this *SuiteInitializeContext) TestInitializeContext() {
	target := this.target()
	context := InitializeContext(target)
	assert.NotNil(this.T(), context)
	assert.Equal(this.T(), &target.Global, context.Global)
	assert.Len(this.T(), context.Element, len(target.Elements))
	element0, ok := context.Element[0].(*initializeElement)
	assert.True(this.T(), ok)
	assert.NotNil(this.T(), element0)
	assert.Equal(this.T(), &target.Global, element0.Global)
	assert.Equal(this.T(), target.Elements[0].Excel, element0.ExcelName)
	assert.Equal(this.T(), target.Elements[0].Sheet, element0.SheetName)
	element1, ok := context.Element[1].(*initializeElement)
	assert.True(this.T(), ok)
	assert.NotNil(this.T(), element1)
	assert.Equal(this.T(), &target.Global, element1.Global)
	assert.Equal(this.T(), target.Elements[1].Excel, element1.ExcelName)
	assert.Equal(this.T(), target.Elements[1].Sheet, element1.SheetName)
}
