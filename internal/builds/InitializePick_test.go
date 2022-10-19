package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializePick(t *testing.T) {
	suite.Run(t, new(SuiteInitializePick))
}

type SuiteInitializePick struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializePick) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializePick) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializePick) target() *Config {
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
			{Excel: testdata.ExcelNameReal, Sheet: testdata.SheetData},
		},
	}
	return target
}

func (this *SuiteInitializePick) TestInitializePick() {
	target := this.target()
	context := InitializeContext(target)

	for _, itor := range context.Element {
		assert.Nil(this.T(), InitializeElement(itor))
	} // for

	assert.Nil(this.T(), InitializePick(context))
	context.Close()
}
