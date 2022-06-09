package core

import (
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskNotes(t *testing.T) {
	ctx := mockTaskNotesContext()
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	err := TaskNotes(ctx)
	assert.Nil(t, err)
	assert.Equal(t, "note0", ctx.Columns[0].Note)
	assert.Equal(t, "note1", ctx.Columns[1].Note)
	assert.Equal(t, "note2", ctx.Columns[2].Note)
	assert.Equal(t, "note3", ctx.Columns[3].Note)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskNotesContext()
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	ctx.Global.LineOfNote = 10
	err = TaskNotes(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)
}

func mockTaskNotesContext() *Context {
	return &Context{
		Global: &Global{
			LineOfNote: 2,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
		Columns: []*Column{
			{Name: "name0", Field: &FieldPkey{}},
			{Name: "name1", Field: &FieldBool{}},
			{Name: "name2", Field: &FieldInt{}},
			{Name: "name3", Field: &FieldText{}},
		},
	}
}
