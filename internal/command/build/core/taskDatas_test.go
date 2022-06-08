package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestTaskDatas(t *testing.T) {
	ctx := mockTaskDatasContext()
	err := TaskDatas(ctx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, ctx.Columns[0].Datas)
	assert.Equal(t, []string{"1", "2", "3"}, ctx.Columns[1].Datas)
	assert.Equal(t, []string{"1", "2", "3"}, ctx.Columns[2].Datas)

	ctx = mockTaskDatasContext()
	ctx.Sheets = Sheets{}
	err = TaskDatas(ctx)
	assert.Nil(t, err)
}

func mockTaskDatasContext() *Context {
	return &Context{
		Progress: util.NewProgress(0, "", ioutil.Discard),
		Global: &Global{
			LineOfData: 3,
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
