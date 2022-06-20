package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskUtil(t *testing.T) {
	task := mockTaskUtil()
	assert.Equal(t, "real.xlsx(Data)", task.originalName())
	assert.Equal(t, "RealData", task.structName())
	assert.Equal(t, "real", task.excelName())
	task.close()

	task = mockTaskUtil()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	assert.True(t, task.sheetExists())
	task.element.Sheet = "?????"
	assert.False(t, task.sheetExists())
	task.close()

	task = mockTaskUtil()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	rows := task.getRows(1)
	assert.NotNil(t, rows)
	_ = rows.Close()
	rows = task.getRows(0)
	assert.Nil(t, rows)
	rows = task.getRows(10)
	assert.Nil(t, rows)
	task.element.Sheet = "?????"
	rows = task.getRows(1)
	assert.Nil(t, rows)
	task.close()

	task = mockTaskUtil()
	task.excel = testdata.GetTestExcel(testdata.Defect1Excel)
	cols := task.getRowContent(1)
	assert.Equal(t, []string{"name0#pkey", "name1#bool", "name2#int", "name3#text"}, cols)
	cols = task.getRowContent(2)
	assert.Equal(t, []string{}, cols)
	cols = task.getRowContent(0)
	assert.Nil(t, cols)
	cols = task.getRowContent(10)
	assert.Nil(t, cols)
	task.element.Sheet = "?????"
	cols = task.getRowContent(1)
	assert.Nil(t, rows)
	task.close()
}

func mockTaskUtil() *Task {
	return &Task{
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
