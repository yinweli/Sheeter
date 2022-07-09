package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestCsReader(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskCsReader()
	err := task.runJsonCsReader()
	assert.Nil(t, err)
	bytes, err := ioutil.ReadFile(task.jsonCsReaderFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskCsReaderString(), string(bytes[:]))
	task.close()

	task = mockTaskCsReader()
	task.element.Excel = "?????.xlsx"
	err = task.runJsonCsReader()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskCsReader()
	task.element.Sheet = "?????"
	err = task.runJsonCsReader()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathJsonCs)
	assert.Nil(t, err)
}

func mockTaskCsReader() *Task {
	return &Task{
		global: &Global{},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
	}
}

func mockTaskCsReaderString() string {
	return `// generated by sheeter, DO NOT EDIT.

namespace sheeter {
    using System;
    using System.Collections.Generic;

    using Newtonsoft.Json;

    public partial class RealDataReader {
		public static readonly string JsonFileName = "realData.json";

        public static Dictionary<string, RealData> FromJson(string data) {
            return JsonConvert.DeserializeObject<Dictionary<string, RealData>>(data);
        }
    }
}`
}
