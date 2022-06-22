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
	assert.Equal(t, "sheeter", task.namespace())
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
	rows, err := task.getRows(1)
	assert.Nil(t, err)
	assert.NotNil(t, rows)
	_ = rows.Close()
	rows, err = task.getRows(10)
	assert.Nil(t, err)
	assert.NotNil(t, rows)
	_ = rows.Close()
	rows, err = task.getRows(0)
	assert.NotNil(t, err)
	task.element.Sheet = "?????"
	rows, err = task.getRows(1)
	assert.NotNil(t, err)
	task.close()

	task = mockTaskUtil()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	cols, err := task.getRowContent(1)
	assert.Nil(t, err)
	assert.Equal(t, []string{"name0#pkey", "name1#bool", "name2#int", "name3#text", "empty#empty"}, cols)
	cols, err = task.getRowContent(2)
	assert.Nil(t, err)
	assert.Equal(t, []string{"note0", "note1", "note2", "note3", "empty"}, cols)
	cols, err = task.getRowContent(10)
	assert.NotNil(t, err)
	cols, err = task.getRowContent(0)
	assert.NotNil(t, err)
	task.element.Sheet = "?????"
	cols, err = task.getRowContent(1)
	assert.NotNil(t, err)
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
