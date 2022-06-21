package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJson(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJson()
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err := task.executeJson()
	assert.Nil(t, err)
	bytes, err := ioutil.ReadFile(task.jsonFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonString(), string(bytes[:]))
	task.close()

	task = mockTaskJson()
	task.excel = testdata.GetTestExcel(testdata.Defect9Excel)
	err = task.executeJson()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskJson()
	task.excel = testdata.GetTestExcel(testdata.Defect10Excel)
	err = task.executeJson()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskJson()
	task.element.Excel = "?????.xlsx"
	task.excel = testdata.GetTestExcel(testdata.RealExcel)
	err = task.executeJson()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathJson)
	assert.Nil(t, err)
}

func TestJsonWrite(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	filePath := "json/test.txt"
	value := map[string]string{"data": "value"}
	jsons, _ := json.MarshalIndent(value, "", "    ")

	err := jsonWrite(value, filePath, true)
	assert.Nil(t, err)

	err = jsonWrite(value, filePath, false)
	assert.Nil(t, err)

	bytes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, jsons, bytes)

	err = jsonWrite(value, "????/????.txt", false)
	assert.NotNil(t, err)

	err = jsonWrite(value, "????.txt", false)
	assert.NotNil(t, err)

	err = os.RemoveAll(path.Dir(filePath))
	assert.Nil(t, err)
}

func mockTaskJson() *Task {
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

func mockTaskJsonString() string {
	return `{
    "1": {
        "name0": 1,
        "name1": true,
        "name2": 1,
        "name3": "a"
    },
    "2": {
        "name0": 2,
        "name1": false,
        "name2": 2,
        "name3": "b"
    },
    "3": {
        "name0": 3,
        "name1": true,
        "name2": 3,
        "name3": "c"
    }
}`
}
