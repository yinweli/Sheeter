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
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	err := TaskJson(ctx)
	assert.Nil(t, err)
	bytes, err := ioutil.ReadFile(ctx.JsonFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonString(), string(bytes[:]))
	util.SilentClose(ctx.Excel)

	ctx = mockTaskJsonContext()
	ctx.Global.LineOfData = 10
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	err = TaskJson(ctx)
	assert.Nil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskJsonContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect9Excel)
	err = TaskJson(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskJsonContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect10Excel)
	err = TaskJson(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskJsonContext()
	ctx.Element.Excel = "?????.xlsx"
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	err = TaskJson(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	err = os.RemoveAll(PathJson)
	assert.Nil(t, err)
}

func mockTaskJsonContext() *Context {
	return &Context{
		Global: &Global{
			LineOfData: 3,
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

func mockTaskJsonString() string {
	return `[
    {
        "name0": 1,
        "name1": true,
        "name2": 1,
        "name3": "a"
    },
    {
        "name0": 2,
        "name1": false,
        "name2": 2,
        "name3": "b"
    },
    {
        "name0": 3,
        "name1": true,
        "name2": 3,
        "name3": "c"
    }
]`
}
