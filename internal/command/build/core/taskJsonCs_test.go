package core

import (
	"io/ioutil"
	"os"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskJsonCs(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	ctx := mockTaskJsonCsContext()
	err := TaskJsonCs(ctx)
	assert.Nil(t, err)
	assert.FileExists(t, ctx.JsonCsFilePath())

	bytes, err := ioutil.ReadFile(ctx.JsonCsFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonCsString(), string(bytes[:]))

	ctx = mockTaskJsonCsContext()
	ctx.Element.Excel = "?????.xlsx"
	err = TaskJsonCs(ctx)
	assert.NotNil(t, err)

	err = os.RemoveAll(PathJsonCs)
	assert.Nil(t, err)
}

func mockTaskJsonCsContext() *Context {
	return &Context{
		Progress: util.NewProgress(0, "", ioutil.Discard),
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldPkey{}},
			{Note: "note1", Name: "name1", Field: &FieldBool{}},
			{Note: "note2", Name: "name2", Field: &FieldText{}},
		},
	}
}

func mockTaskJsonCsString() string {
	return `// generation by sheeter ^o<

using System;
using System.Collections.Generic;

namespace Sheeter {
    public class ExcelSheet { 
        public const string fileName = "excelSheet.json";
        public int Name0; // note0
        public bool Name1; // note1
        public string Name2; // note2
    }
} // namespace Sheeter
`
}
