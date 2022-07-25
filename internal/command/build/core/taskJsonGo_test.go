package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonGo(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJsonGo()
	err := task.runJsonSchema()
	assert.Nil(t, err)
	err = task.runJsonGo()
	assert.Nil(t, err)
	bytes, err := os.ReadFile(task.jsonGoFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonGoString(), string(bytes))
	task.close()

	task = mockTaskJsonGo()
	task.element.Excel = testdata.UnknownStr
	err = task.runJsonGo()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskJsonGo()
	task.element.Sheet = testdata.UnknownStr
	err = task.runJsonGo()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathSchema)
	assert.Nil(t, err)
	err = os.RemoveAll(pathJsonGo)
	assert.Nil(t, err)
}

func mockTaskJsonGo() *Task {
	return &Task{
		global: &Global{},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
		columns: []*Column{
			{Name: "name0", Note: "note0", Field: &FieldPkey{}},
			{Name: "name1", Note: "note1", Field: &FieldBool{}},
			{Name: "name2", Note: "note2", Field: &FieldInt{}},
			{Name: "name3", Note: "note3", Field: &FieldText{}},
		},
	}
}

func mockTaskJsonGoString() string {
	return `package sheeter

type RealData struct {
	Name0 int64  ` + "`json:\"name0\"`" + `
	Name1 bool   ` + "`json:\"name1\"`" + `
	Name2 int64  ` + "`json:\"name2\"`" + `
	Name3 string ` + "`json:\"name3\"`" + `
}
`
}
