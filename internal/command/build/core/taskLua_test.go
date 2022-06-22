package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskLua(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskLua()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err := task.runLua()
	assert.Nil(t, err)
	bytes, err := ioutil.ReadFile(task.luaFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskLuaString(), string(bytes[:]))
	task.close()

	task = mockTaskLua()
	task.excel = testdata.GetTestExcel(testdata.EmptyExcel)
	err = task.runLua()
	assert.Nil(t, err)
	bytes, err = ioutil.ReadFile(task.luaFilePath())
	assert.Nil(t, err)
	assert.Equal(t, "RealData = { \n}", string(bytes[:]))
	task.close()

	task = mockTaskLua()
	task.excel = testdata.GetTestExcel(testdata.Defect9Excel)
	err = task.runLua()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskLua()
	task.element.Excel = "?????.xlsx"
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err = task.runLua()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskLua()
	task.element.Sheet = "?????"
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err = task.runLua()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathLua)
	assert.Nil(t, err)
}

func mockTaskLua() *Task {
	return &Task{
		global: &Global{
			LineOfData: 3,
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
		columns: []*Column{
			{Name: "name0", Field: &FieldPkey{}},
			{Name: "name1", Field: &FieldBool{}},
			{Name: "name2", Field: &FieldInt{}},
			{Name: "name3", Field: &FieldText{}},
		},
	}
}

func mockTaskLuaString() string {
	return `RealData = { 
[1] = { name0 = 1, name1 = true, name2 = 1, name3 = "a",  },
[2] = { name0 = 2, name1 = false, name2 = 2, name3 = "b",  },
[3] = { name0 = 3, name1 = true, name2 = 3, name3 = "c",  },
}`
}