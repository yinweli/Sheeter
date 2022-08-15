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
	workDir string
}

func (this *SuiteTask) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
}

func (this *SuiteTask) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	_ = os.RemoveAll(pathJson)
	_ = os.RemoveAll(pathJsonCs)
	_ = os.RemoveAll(pathJsonGo)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTask) target() *Task {
	target := &Task{
		Path:        testdata.RootPath,
		Bom:         true,
		LineOfField: 1,
		LineOfLayer: 2,
		LineOfNote:  2,
		LineOfData:  4,
		Excel:       testdata.RealExcel,
		Sheet:       testdata.SheetName,
		Progress:    mpb.New(mpb.WithOutput(nil)),
	}
	return target
}

func (this *SuiteTask) TestRun() {
	target := this.target()
	assert.Nil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Path = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.LineOfField = 10
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.LineOfNote = 10
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect1Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect2Excel
	assert.Nil(this.T(), target.Run()) // 測試其實會成功
	target.close()

	target = this.target()
	target.Excel = testdata.Defect3Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect4Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect5Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect6Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect7Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect8Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.Defect9Excel
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run())
	target.close()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.Run())
	target.close()
}

func (this *SuiteTask) TestCheck() {
	target := this.target()
	assert.Nil(this.T(), target.check())

	target = this.target()
	target.LineOfField = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.LineOfLayer = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.LineOfNote = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.LineOfData = 0
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.LineOfField = 4
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.LineOfLayer = 4
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.LineOfNote = 4
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Excel = ""
	assert.NotNil(this.T(), target.check())

	target = this.target()
	target.Sheet = ""
	assert.NotNil(this.T(), target.check())
}
