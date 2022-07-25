package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vbauerster/mpb/v7"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTask(t *testing.T) {
	suite.Run(t, new(SuiteTask))
}

type SuiteTask struct {
	suite.Suite
	workDir  string
	progress *mpb.Progress
}

func (this *SuiteTask) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.progress = mpb.New(mpb.WithOutput(nil))
}

func (this *SuiteTask) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJson)
	_ = os.RemoveAll(pathJsonCs)
	_ = os.RemoveAll(pathJsonGo)
	_ = os.RemoveAll(pathLua)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTask) target() *Task {
	return &Task{
		global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfField: 1,
			LineOfNote:  2,
			LineOfData:  3,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}

func (this *SuiteTask) TestTask() {
	target := this.target()
	defer target.close()

	assert.Nil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcelPath() {
	target := this.target()
	defer target.close()

	target.global.ExcelPath = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskLineOfField() {
	target := this.target()
	defer target.close()

	target.global.LineOfField = 10
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskLineOfNote() {
	target := this.target()
	defer target.close()

	target.global.LineOfNote = 10
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel1() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect1Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel2() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect2Excel
	assert.Nil(this.T(), target.Run(this.progress)) // 測試其實會成功
}

func (this *SuiteTask) TestTaskExcel3() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect3Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel4() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect4Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel5() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect5Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel6() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect6Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel7() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect7Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel8() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect8Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskExcel9() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.Defect9Excel
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskUnknownExcel() {
	target := this.target()
	defer target.close()

	target.element.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestTaskUnknownSheet() {
	target := this.target()
	defer target.close()

	target.element.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run(this.progress))
}

func (this *SuiteTask) TestNewTask() {
	assert.NotNil(this.T(), NewTask(nil, nil))
}
