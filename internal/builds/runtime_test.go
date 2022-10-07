package builds

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/internal/excels"
	"github.com/yinweli/Sheeter/testdata"
)

func TestRuntime(t *testing.T) {
	suite.Run(t, new(SuiteRuntime))
}

type SuiteRuntime struct {
	suite.Suite
	workDir string
}

func (this *SuiteRuntime) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteRuntime) TearDownSuite() {
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteRuntime) target() *RuntimeSector {
	target := &RuntimeSector{
		Element: Element{
			Excel: testdata.ExcelNameReal,
			Sheet: testdata.SheetName,
		},
		excel: &excels.Excel{},
	}
	return target
}

func (this *SuiteRuntime) TestOpenExcel() {
	target := this.target()
	assert.Nil(this.T(), target.OpenExcel())
	target.CloseExcel()

	target = this.target()
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.OpenExcel())

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.OpenExcel())
}

func (this *SuiteRuntime) TestGetExcelLine() {
	target := this.target()
	assert.Nil(this.T(), target.OpenExcel())
	line, err := target.GetExcelLine(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), line)
	target.CloseExcel()

	target = this.target()
	assert.Nil(this.T(), target.OpenExcel())
	target.Sheet = testdata.UnknownStr
	_, err = target.GetExcelLine(1)
	assert.NotNil(this.T(), err)
	target.CloseExcel()
}

func (this *SuiteRuntime) TestGetExcelData() {
	target := this.target()
	assert.Nil(this.T(), target.OpenExcel())
	data, err := target.GetExcelData(1)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), data)
	target.CloseExcel()

	target = this.target()
	assert.Nil(this.T(), target.OpenExcel())
	target.Sheet = testdata.UnknownStr
	_, err = target.GetExcelData(1)
	assert.NotNil(this.T(), err)
	target.CloseExcel()
}
