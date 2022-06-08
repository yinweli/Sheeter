package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskExcel(t *testing.T) {
	ctx := mockTaskExcelContext()
	err := TaskExcel(ctx, ioutil.Discard)
	assert.Nil(t, err)
	assert.NotNil(t, ctx.Progress)
	assert.NotNil(t, ctx.Sheets)
	assert.Equal(t, 12, len(ctx.Sheets))
	assert.Equal(t, 16, len(ctx.Sheets[0]))
	assert.Equal(t, "checkpoint", ctx.Sheets[0][15])
	assert.Equal(t, "checkpoint", ctx.Sheets[11][15])

	ctx = mockTaskExcelContext()
	ctx.Global.ExcelPath = ""
	err = TaskExcel(ctx, ioutil.Discard)
	assert.NotNil(t, err)

	ctx = mockTaskExcelContext()
	ctx.Element.Excel = testdata.FakeExcel
	err = TaskExcel(ctx, ioutil.Discard)
	assert.NotNil(t, err)

	ctx = mockTaskExcelContext()
	ctx.Element.Excel = testdata.DefectExcel
	ctx.Element.Sheet = testdata.DefectSheet
	err = TaskExcel(ctx, ioutil.Discard)
	assert.NotNil(t, err)

	ctx = mockTaskExcelContext()
	ctx.Element.Excel = "?????"
	err = TaskExcel(ctx, ioutil.Discard)
	assert.NotNil(t, err)

	ctx = mockTaskExcelContext()
	ctx.Element.Sheet = "?????"
	err = TaskExcel(ctx, ioutil.Discard)
	assert.NotNil(t, err)
}

func mockTaskExcelContext() *Context {
	return &Context{
		Global: &Global{
			ExcelPath: testdata.RootPath,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
	}
}
