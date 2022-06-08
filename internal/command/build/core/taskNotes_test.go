package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestTaskNotes(t *testing.T) {
	ctx := mockTaskNotesContext()
	err := TaskNotes(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "note0", ctx.Columns[0].Note)
	assert.Equal(t, "note1", ctx.Columns[1].Note)
	assert.Equal(t, "note2", ctx.Columns[2].Note)

	ctx = mockTaskNotesContext()
	ctx.Sheets = Sheets{}
	err = TaskNotes(ctx)
	assert.NotNil(t, err)
}

func mockTaskNotesContext() *Context {
	return &Context{
		Progress: util.NewProgress(0, "", ioutil.Discard),
		Global: &Global{
			LineOfNote: 2,
		},
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Sheets: Sheets{
			{"name0#pkey", "name1#int", "name2#int"},
			{"note0", "note1", "note2"},
			{"1", "1", "1"},
			{"2", "2", "2"},
			{"3", "3", "3"},
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldPkey{}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}},
		},
	}
}
