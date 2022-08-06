package tasks

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskLua(t *testing.T) {
	suite.Run(t, new(SuiteTaskLua))
}

type SuiteTaskLua struct {
	suite.Suite
	workDir    string
	dataBytes  []byte
	emptyBytes []byte
}

func (this *SuiteTaskLua) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.dataBytes = []byte(`-- generated by sheeter, DO NOT EDIT.

RealData = { 
[1] = { name0 = 1, name1 = true, name2 = 1, name3 = "a",  },
[2] = { name0 = 2, name1 = false, name2 = 2, name3 = "b",  },
[3] = { name0 = 3, name1 = true, name2 = 3, name3 = "c",  },
}`)
	this.emptyBytes = []byte(`-- generated by sheeter, DO NOT EDIT.

RealData = { 
}`)
}

func (this *SuiteTaskLua) TearDownSuite() {
	_ = os.RemoveAll(pathLua)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTaskLua) target() *Task {
	target := NewTask(nil, nil)
	target.global = &Global{
		LineOfData: 4,
	}
	target.element = &Element{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	target.columns = []*Column{
		{Name: "name0", Field: &fields.FieldPkey{}},
		{Name: "name1", Field: &fields.FieldBool{}},
		{Name: "name2", Field: &fields.FieldInt{}},
		{Name: "name3", Field: &fields.FieldText{}},
	}
	return target
}

func (this *SuiteTaskLua) TestTaskLua() {
	target := this.target()
	target.excel = testdata.GetTestExcel(testdata.RealExcel)
	assert.Nil(this.T(), target.runLua())
	testdata.CompareFile(this.T(), target.luaFilePath(), this.dataBytes)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.EmptyExcel)
	assert.Nil(this.T(), target.runLua())
	testdata.CompareFile(this.T(), target.luaFilePath(), this.emptyBytes)
	target.close()

	target = this.target()
	target.excel = testdata.GetTestExcel(testdata.Defect9Excel)
	assert.NotNil(this.T(), target.runLua())
	target.close()

	target = this.target()
	target.element.Excel = testdata.UnknownStr
	target.excel = testdata.GetTestExcel(testdata.RealExcel)
	assert.NotNil(this.T(), target.runLua())
	target.close()

	target = this.target()
	target.element.Sheet = testdata.UnknownStr
	target.excel = testdata.GetTestExcel(testdata.RealExcel)
	assert.NotNil(this.T(), target.runLua())
	target.close()
}