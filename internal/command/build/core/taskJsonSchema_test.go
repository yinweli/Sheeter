package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonSchema(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJsonSchema()
	err := task.runJsonSchema()
	assert.Nil(t, err)
	bytes, err := os.ReadFile(task.jsonSchemaFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonSchemaString(), string(bytes))
	task.close()

	err = os.RemoveAll(pathSchema)
	assert.Nil(t, err)
}

func mockTaskJsonSchema() *Task {
	return &Task{
		global: &Global{},
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

func mockTaskJsonSchemaString() string {
	return `{
    "name0": 0,
    "name1": false,
    "name2": 0,
    "name3": ""
}`
}
