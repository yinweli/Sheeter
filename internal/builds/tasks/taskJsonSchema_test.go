package tasks

import (
	"os"
	"testing"

	"github.com/yinweli/Sheeter/internal/build/fields"
	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
	target := &Task{
		Excel: testdata.RealExcel,
		Sheet: testdata.SheetName,
		columns: []*Column{
			{Name: "name0", Field: &fields.Pkey{}},
			{Name: "name1", Field: &fields.Bool{}},
			{Name: "name2", Field: &fields.Int{}},
			{Name: "name3", Field: &fields.Text{}},
		},
	}
	return target
}

func (this *SuiteTaskJsonSchema) TestJsonSchema() {
	target := this.target()
	assert.Nil(this.T(), target.jsonSchema())
	testdata.CompareFile(this.T(), target.jsonSchemaFilePath(), this.dataBytes)
	target.close()
}
