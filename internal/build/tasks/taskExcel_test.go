package tasks

import (
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestTaskExcel(t *testing.T) {
	suite.Run(t, new(SuiteTaskExcel))
}

type SuiteTaskExcel struct {
	suite.Suite
}

func (this *SuiteTaskExcel) target() *Task {
	target := NewTask(nil, nil)
	target.global = &Global{
		ExcelPath: testdata.RootPath,
	}
	target.element = &Element{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	return target
}

func (this *SuiteTaskExcel) TestTaskExcel() {
	target := this.target()
	assert.Nil(this.T(), target.runExcel())
	assert.NotNil(this.T(), target.excel)
	target.close()

	target = this.target()
	target.global.ExcelPath = ""
	assert.NotNil(this.T(), target.runExcel())
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect1Excel
	assert.NotNil(this.T(), target.runExcel())
	target.close()

	target = this.target()
	target.element.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.runExcel())
	target.close()

	target = this.target()
	target.element.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.runExcel())
	target.close()
}
