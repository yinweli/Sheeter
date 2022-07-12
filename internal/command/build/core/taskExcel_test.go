package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskExcel(t *testing.T) {
	task := mockTaskExcel()
	err := task.runExcel()
	assert.Nil(t, err)
	assert.NotNil(t, task.excel)
	task.close()

	task = mockTaskExcel()
	task.global.ExcelPath = ""
	err = task.runExcel()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskExcel()
	task.element.Excel = testdata.Defect1Excel
	err = task.runExcel()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskExcel()
	task.element.Excel = testdata.UnknownStr
	err = task.runExcel()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskExcel()
	task.element.Sheet = testdata.UnknownStr
	err = task.runExcel()
	assert.NotNil(t, err)
	task.close()
}

func mockTaskExcel() *Task {
	return &Task{
		global: &Global{
			ExcelPath: testdata.RootPath,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
