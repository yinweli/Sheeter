package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestTaskColumns(t *testing.T) {
	ctx := mockTaskColumnsContext()
	err := TaskColumns(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(ctx.Columns))
	assert.Equal(t, "name0", ctx.Columns[0].Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), ctx.Columns[0].Field.TypeExcel())
	assert.Equal(t, "name1", ctx.Columns[1].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), ctx.Columns[1].Field.TypeExcel())
	assert.Equal(t, "name2", ctx.Columns[2].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), ctx.Columns[2].Field.TypeExcel())
	assert.NotNil(t, ctx.Pkey)
	assert.Equal(t, "name0", ctx.Pkey.Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), ctx.Pkey.Field.TypeExcel())

	ctx = mockTaskColumnsContext()
	ctx.Sheets = Sheets{}
	err = TaskColumns(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskColumnsContext()
	ctx.Sheets[ctx.Global.GetLineOfField()] = []string{}
	err = TaskColumns(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskColumnsContext()
	ctx.Sheets[ctx.Global.GetLineOfField()] = []string{"", "", ""}
	err = TaskColumns(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskColumnsContext()
	ctx.Sheets[ctx.Global.GetLineOfField()] = []string{"name0#????", "name1#bool", "name2#int"}
	err = TaskColumns(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskColumnsContext()
	ctx.Sheets[ctx.Global.GetLineOfField()] = []string{"name0#pkey", "name1#pkey", "name2#int"}
	err = TaskColumns(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskColumnsContext()
	ctx.Sheets[ctx.Global.GetLineOfField()] = []string{"name0#int", "name1#int", "name2#int"}
	err = TaskColumns(ctx)
	assert.NotNil(t, err)
}

func mockTaskColumnsContext() *Context {
	return &Context{
		Progress: util.NewProgress(0, "", ioutil.Discard),
		Global: &Global{
			LineOfField: 1,
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
	}
}
