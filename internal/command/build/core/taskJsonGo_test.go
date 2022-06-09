package core

import (
	"io/ioutil"
	"os"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestTaskJsonGo(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	ctx := mockTaskJsonGoContext()
	err := TaskJsonGo(ctx)
	assert.Nil(t, err)
	assert.FileExists(t, ctx.JsonGoFilePath())

	bytes, err := ioutil.ReadFile(ctx.JsonGoFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonGoString(), string(bytes[:]))

	ctx = mockTaskJsonGoContext()
	ctx.Element.Excel = "?????.xlsx"
	err = TaskJsonGo(ctx)
	assert.NotNil(t, err)

	err = os.RemoveAll(PathJsonGo)
	assert.Nil(t, err)
}

func mockTaskJsonGoContext() *Context {
	return &Context{
		// Progress: util.NewProgress(0, "", ioutil.Discard),
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Columns: []*Column{
			{Name: "name0", Note: "note0", Field: &FieldPkey{}},
			{Name: "name1", Note: "note1", Field: &FieldBool{}},
			{Name: "name2", Note: "note2", Field: &FieldText{}},
		},
	}
}

func mockTaskJsonGoString() string {
	return `// generation by sheeter ^o<

package sheeter

const ExcelSheetFileName = "excelSheet.json" // json file name

type ExcelSheet struct {
	Name0 int32  // note0
	Name1 bool   // note1
	Name2 string // note2
}
`
}
