package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/testdata"
)

func TestInitializeSector(t *testing.T) {
	suite.Run(t, new(SuiteInitializeSector))
}

type SuiteInitializeSector struct {
	suite.Suite
	workDir string
}

func (this *SuiteInitializeSector) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteInitializeSector) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteInitializeSector) target() *RuntimeSector {
	target := &RuntimeSector{
		Global: Global{
			LineOfField: 1,
			LineOfLayer: 2,
			LineOfNote:  3,
		},
		Element: Element{
			Excel: testdata.ExcelNameReal,
			Sheet: testdata.SheetName,
		},
	}
	return target
}

func (this *SuiteInitializeSector) TestInitializeSector() {
	target := this.target()
	assert.Nil(this.T(), initializeSector(target))
	assert.NotNil(this.T(), target.Mixed)
	assert.NotNil(this.T(), target.excel)
	assert.NotNil(this.T(), target.layoutJson)
	assert.NotNil(this.T(), target.layoutType)
	assert.NotNil(this.T(), target.layoutDepend)
	target.CloseExcel()

	target = this.target()
	target.LineOfField = 10
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.LineOfLayer = 10
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.LineOfNote = 10
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidFile
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameCleanAll
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameCleanField
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidField
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidLayer
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidLayout
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidPkeyZero
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.ExcelNameInvalidPkeyDupl
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), initializeSector(target))
	target.CloseExcel()
}
