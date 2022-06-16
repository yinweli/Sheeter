package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
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
		global: &Global{},
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
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Text;

namespace Sheeter {
    public class RealData { 
        public const string fileName = "realData.json";
        public int name0; // note0
        public bool name1; // note1
        public int name2; // note2
        public string name3; // note3

        public static Dictionary<int, RealData> Parse(string s) {
            return JsonConvert.DeserializeObject<Dictionary<int, RealData>>(s);
        }

        public static Dictionary<int, RealData> Parse(byte[] b)
        {
            return Parse(Encoding.UTF8.GetString(b));
        }
    }
} // namespace Sheeter
`
}
