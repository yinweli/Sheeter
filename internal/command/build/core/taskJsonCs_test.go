package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/yinweli/Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskJsonCs(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJsonCs()
	err := task.executeJsonCs()
	assert.Nil(t, err)
	bytes, err := ioutil.ReadFile(task.jsonCsFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonCsString(), string(bytes[:]))
	task.close()

	task = mockTaskJsonCs()
	task.element.Excel = "?????.xlsx"
	err = task.executeJsonCs()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathJsonCs)
	assert.Nil(t, err)
}

func mockTaskJsonCs() *Task {
	return &Task{
		global: &Global{
			CppLibraryPath: "nlohmann/json.hpp",
		},
		element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.SheetName,
		},
		columns: []*Column{
			{Name: "name0", Note: "note0", Field: &FieldPkey{}},
			{Name: "name1", Note: "note1", Field: &FieldBool{}},
			{Name: "name2", Note: "note2", Field: &FieldInt{}},
			{Name: "name3", Note: "note3", Field: &FieldText{}},
		},
	}
}

func mockTaskJsonCsString() string {
	return `// generation by sheeter ^o<

using System;
using System.Collections.Generic;

namespace Sheeter {
    public class RealData { 
        public const string fileName = "realData.json";
        public int Name0; // note0
        public bool Name1; // note1
        public int Name2; // note2
        public string Name3; // note3
    }
} // namespace Sheeter
`
}
