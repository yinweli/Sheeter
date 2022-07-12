package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonCs(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJsonCs()
	err := task.runJsonSchema()
	assert.Nil(t, err)
	err = task.runJsonCs()
	assert.Nil(t, err)
	bytes, err := os.ReadFile(task.jsonCsFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonCsString(), string(bytes))
	task.close()

	task = mockTaskJsonCs()
	task.element.Excel = testdata.UnknownExcel
	err = task.runJsonCs()
	assert.NotNil(t, err)
	task.close()

	task = mockTaskJsonCs()
	task.element.Sheet = testdata.UnknownStr
	err = task.runJsonCs()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathSchema)
	assert.Nil(t, err)
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
	return `namespace sheeter
{
    using System;
    using System.Collections.Generic;

    using System.Globalization;
    using Newtonsoft.Json;
    using Newtonsoft.Json.Converters;

    public partial class RealData
    {
        [JsonProperty("name0")]
        public long Name0 { get; set; }

        [JsonProperty("name1")]
        public bool Name1 { get; set; }

        [JsonProperty("name2")]
        public long Name2 { get; set; }

        [JsonProperty("name3")]
        public string Name3 { get; set; }
    }
}
`
}
