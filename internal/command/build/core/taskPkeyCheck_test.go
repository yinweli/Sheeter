package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestTaskPkeyCheck(t *testing.T) {
	ctx := mockTaskPkeyCheckContext()
	err := TaskPkeyCheck(ctx)
	assert.Nil(t, err)

	ctx = mockTaskPkeyCheckContext()
	ctx.Pkey.Datas = []string{"1", "2", "3", "3"}
	err = TaskPkeyCheck(ctx)
	assert.NotNil(t, err)
}

func mockTaskPkeyCheckContext() *Context {
	return &Context{
		Progress: util.NewProgress(0, "", ioutil.Discard),
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Pkey: &Column{
			Note:  "note",
			Name:  "name",
			Field: &FieldPkey{},
			Datas: []string{"1", "2", "3"},
		},
	}
}
