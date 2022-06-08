package core

import (
	"io/ioutil"
	"os"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskJson(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	ctx := mockTaskJsonContext()
	err := TaskJson(ctx)
	assert.Nil(t, err)
	assert.FileExists(t, ctx.JsonFilePath())

	bytes, err := ioutil.ReadFile(ctx.JsonFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonString(), string(bytes[:]))

	ctx = mockTaskJsonContext()
	ctx.Columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"x", "2", "3"}},
	}
	err = TaskJson(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskJsonContext()
	ctx.Element.Excel = "?????.xlsx"
	err = TaskJson(ctx)
	assert.NotNil(t, err)

	err = os.RemoveAll(PathJson)
	assert.Nil(t, err)
}

func mockTaskJsonContext() *Context {
	return &Context{
		Progress: util.NewProgress(0, "", ioutil.Discard),
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldPkey{}, Datas: []string{"1", "2", "3"}},
			{Note: "note1", Name: "name1", Field: &FieldBool{}, Datas: []string{"false", "true", "false"}},
			{Note: "note2", Name: "name2", Field: &FieldText{}, Datas: []string{"text1", "text2", "text3"}},
		},
	}
}

func mockTaskJsonString() string {
	return `[
    {
        "name0": 1,
        "name1": false,
        "name2": "text1"
    },
    {
        "name0": 2,
        "name1": true,
        "name2": "text2"
    },
    {
        "name0": 3,
        "name1": false,
        "name2": "text3"
    }
]`
}
