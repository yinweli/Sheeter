package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTask(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTask()
	err := task.Execute()
	assert.Nil(t, err)
	task.close()

	task = mockTask()
	task.global.ExcelPath = "?????"
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.global.LineOfField = 10
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.global.LineOfNote = 10
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect1Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect2Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect3Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect4Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect5Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect6Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect7Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect8Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect9Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect10Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect11Excel
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = "?????.xlsx"
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Sheet = "?????"
	err = task.Execute()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathJson)
	assert.Nil(t, err)
	err = os.RemoveAll(pathJsonCpp)
	assert.Nil(t, err)
	err = os.RemoveAll(pathJsonCs)
	assert.Nil(t, err)
	err = os.RemoveAll(pathJsonGo)
	assert.Nil(t, err)
}

func TestNewTask(t *testing.T) {
	task := NewTask(nil, nil, nil)
	assert.NotNil(t, task)
}

func mockTask() *Task {
	return &Task{
		global: &Global{
			ExcelPath:      testdata.RootPath,
			CppLibraryPath: "nlohmann/json.hpp",
			LineOfField:    1,
			LineOfNote:     2,
			LineOfData:     3,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
