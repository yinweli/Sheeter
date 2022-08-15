package tasks

import (
	"os"
	"testing"

	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestTaskJson(t *testing.T) {
	suite.Run(t, new(SuiteTaskJson))
}

type SuiteTaskJson struct {
	suite.Suite
	workDir    string
	dataBytes  []byte
	emptyBytes []byte
}

func (this *SuiteTaskJson) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.dataBytes = []byte(`{
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
	target := &Task{
		LineOfData: 4,
		Excel:      testdata.RealExcel,
		Sheet:      testdata.SheetName,
		columns: []*Column{
			{Name: "name0", Field: &fields.Pkey{}},
			{Name: "name1", Field: &fields.Bool{}},
			{Name: "name2", Field: &fields.Int{}},
			{Name: "name3", Field: &fields.Text{}},
		},
	}
	return target
}

func (this *SuiteTaskJson) TestJson() {
	target := this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)
	assert.Nil(this.T(), target.json())
	testdata.CompareFile(this.T(), target.jsonFilePath(), this.dataBytes)
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.EmptyExcel)
	assert.Nil(this.T(), target.json())
	testdata.CompareFile(this.T(), target.jsonFilePath(), this.emptyBytes)
	target.close()

	target = this.target()
	target.xlsfile = testdata.GetTestExcel(testdata.Defect9Excel)
	assert.NotNil(this.T(), target.json())
	target.close()

	target = this.target()
	target.Excel = testdata.UnknownStr
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)
	assert.NotNil(this.T(), target.json())
	target.close()

	target = this.target()
	target.Sheet = testdata.UnknownStr
	target.xlsfile = testdata.GetTestExcel(testdata.RealExcel)
	assert.NotNil(this.T(), target.json())
	target.close()
}
