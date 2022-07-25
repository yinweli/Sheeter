package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskExcel(t *testing.T) {
	suite.Run(t, new(SuiteTaskExcel))
}

type SuiteTaskExcel struct {
	suite.Suite
}

func (this *SuiteTaskExcel) target() *Task {
	return &Task{
		global: &Global{
			ExcelPath: testdata.RootPath,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}

func (this *SuiteTaskExcel) TestTaskExcel() {
	target := this.target()
	defer target.close()

	assert.Nil(this.T(), target.runExcel())
	assert.NotNil(this.T(), target.excel)
}

func (this *SuiteTaskExcel) TestTaskExcelPath() {
	target := this.target()
	defer target.close()

	target.global.ExcelPath = ""
	assert.NotNil(this.T(), target.runExcel())
}

func (this *SuiteTaskExcel) TestTaskExcel1() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect1Excel
	assert.NotNil(this.T(), target.runExcel())
}

func (this *SuiteTaskExcel) TestTaskExcelUnknownExcel() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.runExcel())
}

func (this *SuiteTaskExcel) TestTaskExcelUnknownSheet() {
	target := this.target()
	defer target.close()

	target.element.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.runExcel())
}
