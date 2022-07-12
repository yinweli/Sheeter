package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vbauerster/mpb/v7"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTask(t *testing.T) {
	progress := mpb.New(mpb.WithOutput(nil))
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTask()
	err := task.Run(progress)
	assert.Nil(t, err)
	task.close()

	task = mockTask()
	task.global.ExcelPath = testdata.UnknownStr
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.global.LineOfField = 10
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.global.LineOfNote = 10
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect1Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect2Excel // 測試其實會成功
	err = task.Run(progress)
	assert.Nil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect3Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect4Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect5Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect6Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect7Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect8Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.Defect9Excel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Excel = testdata.UnknownExcel
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	task = mockTask()
	task.element.Sheet = testdata.UnknownStr
	err = task.Run(progress)
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathSchema)
	assert.Nil(t, err)
	err = os.RemoveAll(pathJson)
	assert.Nil(t, err)
	err = os.RemoveAll(pathJsonCs)
	assert.Nil(t, err)
	err = os.RemoveAll(pathJsonGo)
	assert.Nil(t, err)
	err = os.RemoveAll(pathLua)
	assert.Nil(t, err)
}

func TestNewTask(t *testing.T) {
	task := NewTask(nil, nil)
	assert.NotNil(t, task)
}

func mockTask() *Task {
	return &Task{
		global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfField: 1,
			LineOfNote:  2,
			LineOfData:  3,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}
