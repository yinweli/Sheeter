package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJson(t *testing.T) {
	suite.Run(t, new(SuiteTaskJson))
}

type SuiteTaskJson struct {
	suite.Suite
	workDir    string
	jsonBytes  []byte
	emptyBytes []byte
}

func (this *SuiteTaskJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.jsonBytes = []byte(`{
    "1": {
        "name0": 1,
        "name1": true,
        "name2": 1,
        "name3": "a"
    },
    "2": {
        "name0": 2,
        "name1": false,
        "name2": 2,
        "name3": "b"
    },
    "3": {
        "name0": 3,
        "name1": true,
        "name2": 3,
        "name3": "c"
    }
}`)
	this.emptyBytes = []byte("{}")
}

func (this *SuiteTaskJson) TearDownSuite() {
	_ = os.RemoveAll(pathJson)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTaskJson) target() *Task {
	return &Task{
		global: &Global{
			LineOfData: 3,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
		columns: []*Column{
			{Name: "name0", Field: &FieldPkey{}},
			{Name: "name1", Field: &FieldBool{}},
			{Name: "name2", Field: &FieldInt{}},
			{Name: "name3", Field: &FieldText{}},
		},
	}
}

func (this *SuiteTaskJson) check(filepath string, expected []byte) {
	actual, err := os.ReadFile(filepath)
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), expected, actual)
}

func (this *SuiteTaskJson) TestTaskJson() {
	target := this.target()
	defer target.close()

	target.excel = testdata.GetTestExcel(testdata.RealExcel)
	assert.Nil(this.T(), target.runJson())
	this.check(target.jsonFilePath(), this.jsonBytes)
}

func (this *SuiteTaskJson) TestTaskJsonExcelEmpty() {
	target := this.target()
	defer target.close()

	target.excel = testdata.GetTestExcel(testdata.EmptyExcel)
	assert.Nil(this.T(), target.runJson())
	this.check(target.jsonFilePath(), this.emptyBytes)
}

func (this *SuiteTaskJson) TestTaskJsonExcel9() {
	target := this.target()
	defer target.close()

	target.excel = testdata.GetTestExcel(testdata.Defect9Excel)
	assert.NotNil(this.T(), target.runJson())
}

func (this *SuiteTaskJson) TestTaskJsonUnknownExcel() {
	target := this.target()
	defer target.close()

	target.excel = testdata.GetTestExcel(testdata.RealExcel)
	target.element.Excel = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJson())
}

func (this *SuiteTaskJson) TestTaskJsonUnknownSheet() {
	target := this.target()
	defer target.close()

	target.excel = testdata.GetTestExcel(testdata.RealExcel)
	target.element.Sheet = testdata.UnknownStr
	assert.NotNil(this.T(), target.runJson())
}
