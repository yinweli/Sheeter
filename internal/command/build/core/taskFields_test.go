package core

import (
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskFields(t *testing.T) {
	ctx := mockTaskFieldsContext()
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	err := TaskFields(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(ctx.Columns))
	assert.Equal(t, "name0", ctx.Columns[0].Name)
	assert.Equal(t, (&FieldPkey{}).TypeExcel(), ctx.Columns[0].Field.TypeExcel())
	assert.Equal(t, "name1", ctx.Columns[1].Name)
	assert.Equal(t, (&FieldBool{}).TypeExcel(), ctx.Columns[1].Field.TypeExcel())
	assert.Equal(t, "name2", ctx.Columns[2].Name)
	assert.Equal(t, (&FieldInt{}).TypeExcel(), ctx.Columns[2].Field.TypeExcel())
	assert.Equal(t, "name3", ctx.Columns[3].Name)
	assert.Equal(t, (&FieldText{}).TypeExcel(), ctx.Columns[3].Field.TypeExcel())
	util.SilentClose(ctx.Excel)

	ctx = mockTaskFieldsContext()
	ctx.Global.LineOfField = 10
	ctx.Excel = testdata.GetTestExcel(testdata.RealExcel)
	err = TaskFields(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskFieldsContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect3Excel)
	err = TaskFields(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskFieldsContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect4Excel)
	err = TaskFields(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskFieldsContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect5Excel)
	err = TaskFields(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskFieldsContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect6Excel)
	err = TaskFields(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskFieldsContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect7Excel)
	err = TaskFields(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)

	ctx = mockTaskFieldsContext()
	ctx.Excel = testdata.GetTestExcel(testdata.Defect8Excel)
	err = TaskFields(ctx)
	assert.NotNil(t, err)
	util.SilentClose(ctx.Excel)
}

func mockTaskFieldsContext() *Context {
	return &Context{
		Global: &Global{
			LineOfField: 1,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
