package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonCs(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJsonCs()
	task.element.Excel = "?????.xlsx"
	err := task.executeJsonCs()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathJsonCs)
	assert.Nil(t, err)
}

func mockTaskJsonCs() *Task {
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
