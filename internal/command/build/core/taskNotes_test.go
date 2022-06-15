package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskNotes(t *testing.T) {
	task := mockTaskNotes()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err := task.executeNotes()
	assert.Nil(t, err)
	assert.Equal(t, "note0", task.columns[0].Note)
	assert.Equal(t, "note1", task.columns[1].Note)
	assert.Equal(t, "note2", task.columns[2].Note)
	assert.Equal(t, "note3", task.columns[3].Note)
	task.close()

	task = mockTaskNotes()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	task.global.LineOfNote = 10
	err = task.executeNotes()
	assert.NotNil(t, err)
	task.close()
}

func mockTaskNotes() *Task {
	return &Task{
		global: &Global{
			LineOfNote: 2,
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
