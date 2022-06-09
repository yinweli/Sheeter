package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskExcel(t *testing.T) {
	ctx := mockTaskExcelContext()
	err := TaskExcel(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, ctx.Excel)

	ctx = mockTaskExcelContext()
	ctx.Global.ExcelPath = ""
	err = TaskExcel(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskExcelContext()
	ctx.Element.Excel = testdata.Defect2Excel
	err = TaskExcel(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskExcelContext()
	ctx.Element.Excel = "?????"
	err = TaskExcel(ctx)
	assert.NotNil(t, err)

	ctx = mockTaskExcelContext()
	ctx.Element.Sheet = "?????"
	err = TaskExcel(ctx)
	assert.NotNil(t, err)
}

func mockTaskExcelContext() *Context {
	return &Context{
		Global: &Global{
			ExcelPath: testdata.RootPath,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
