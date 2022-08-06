package tasks

import (
	"os"
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vbauerster/mpb/v7"
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
	target := NewTask(nil, nil)
	target.global = &Global{
		ExcelPath:   testdata.RootPath,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  2,
		LineOfData:  4,
	}
	target.element = &Element{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	return target
}

func (this *SuiteTask) TestTask() {
	target := this.target()
	assert.Nil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.global.ExcelPath = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.global.LineOfField = 10
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.global.LineOfNote = 10
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect1Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect2Excel
	assert.Nil(this.T(), target.Run(this.progress)) // 測試其實會成功
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect3Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect4Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect5Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect6Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect7Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect8Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.Defect9Excel
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()

	target = this.target()
	target.element.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run(this.progress))
	target.close()
}

func (this *SuiteTask) TestNewTask() {
	assert.NotNil(this.T(), NewTask(nil, nil))
}

func TestGlobal(t *testing.T) {
	suite.Run(t, new(SuiteGlobal))
}

type SuiteGlobal struct {
	suite.Suite
}

func (this *SuiteGlobal) target() *Global {
	return &Global{
		ExcelPath:   "excel",
		Bom:         true,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  3,
		LineOfData:  4,
	}
}

func (this *SuiteGlobal) TestCheck() {
	target := this.target()
	assert.Nil(this.T(), target.Check())

	target = this.target()
	target.LineOfField = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfLayer = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfNote = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfData = 0
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfField = 4
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfLayer = 4
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.LineOfNote = 4
	assert.NotNil(this.T(), target.Check())
}

func TestElement(t *testing.T) {
	suite.Run(t, new(SuiteElement))
}

type SuiteElement struct {
	suite.Suite
}

func (this *SuiteElement) target() *Element {
	return &Element{
		Excel: "excel.xlsx",
		Sheet: "sheet",
	}
}

func (this *SuiteElement) TestCheck() {
	target := this.target()
	assert.Nil(this.T(), target.Check())

	target = this.target()
	target.Excel = ""
	assert.NotNil(this.T(), target.Check())

	target = this.target()
	target.Sheet = ""
	assert.NotNil(this.T(), target.Check())
}
