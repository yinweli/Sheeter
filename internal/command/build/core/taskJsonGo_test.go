package core

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTaskJsonGo(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	task := mockTaskJsonGo()
	err := task.executeJsonGo()
	assert.Nil(t, err)
	bytes, err := ioutil.ReadFile(task.jsonGoFilePath())
	assert.Nil(t, err)
	assert.Equal(t, mockTaskJsonGoString(), string(bytes[:]))
	task.close()

	task = mockTaskJsonGo()
	task.element.Excel = "?????.xlsx"
	err = task.executeJsonGo()
	assert.NotNil(t, err)
	task.close()

	err = os.RemoveAll(pathJsonGo)
	assert.Nil(t, err)
}

func mockTaskJsonGo() *Task {
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

func mockTaskJsonGoString() string {
	return `// generation by sheeter ^o<

package sheeter

const RealDataFileName = "realData.json" // json file name

type RealData struct {
	Name0 int32  ` + "`json:\"name0\"`" + ` // note0
	Name1 bool   ` + "`json:\"name1\"`" + ` // note1
	Name2 int32  ` + "`json:\"name2\"`" + ` // note2
	Name3 string ` + "`json:\"name3\"`" + ` // note3
}
`
}
