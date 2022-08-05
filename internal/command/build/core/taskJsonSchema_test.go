package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonSchema(t *testing.T) {
	suite.Run(t, new(SuiteTaskJsonSchema))
}

type SuiteTaskJsonSchema struct {
	suite.Suite
	workDir   string
	dataBytes []byte
}

func (this *SuiteTaskJsonSchema) SetupSuite() {
	this.workDir = testdata.ChangeWorkDir()
	this.dataBytes = []byte(`{
    "name0": 0,
    "name1": false,
    "name2": 0,
    "name3": ""
}`)
}

func (this *SuiteTaskJsonSchema) TearDownSuite() {
	_ = os.RemoveAll(pathSchema)
	testdata.RestoreWorkDir(this.workDir)
}

func (this *SuiteTaskJsonSchema) target() *Task {
	target := NewTask(nil, nil)
	target.global = &Global{}
	target.element = &Element{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
	}
	target.columns = []*Column{
		{Name: "name0", Field: &FieldPkey{}},
		{Name: "name1", Field: &FieldBool{}},
		{Name: "name2", Field: &FieldInt{}},
		{Name: "name3", Field: &FieldText{}},
	}
	return target
}

func (this *SuiteTaskJsonSchema) TestTaskJsonSchema() {
	target := this.target()
	assert.Nil(this.T(), target.runJsonSchema())
	testdata.CompareFile(this.T(), target.jsonSchemaFilePath(), this.dataBytes)
	target.close()
}
