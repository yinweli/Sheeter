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
	task.element.Excel = "?????.xlsx"
	err := task.executeJsonGo()
	assert.NotNil(t, err)
	task.close()

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
