package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteGo(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	cargo := mockWriteGoCargo()
	path, err := WriteGo(cargo)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(OutputPathGo, "realData.go"), path)
	assert.FileExists(t, path)

	bytes, err := ioutil.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, mockWriteGoString(), string(bytes[:]))

	err = os.RemoveAll(OutputPathGo)
	assert.Nil(t, err)
}

func mockWriteGoCargo() *Cargo {
	return &Cargo{
		Progress: NewProgress(0, "test", ioutil.Discard),
		Element: &Element{
			Excel: "real.xlsx",
			Sheet: "data",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}},
			{Note: "note1", Name: "name1", Field: &FieldBool{}},
			{Note: "note2", Name: "name2", Field: &FieldText{}},
		},
	}
}

func mockWriteGoString() string {
	return `// generation by sheeter ^o<

package sheeter

const RealDataFileName string = "realData.json" // json file name

type RealData struct {
	Name0 int32  // note0
	Name1 bool   // note1
	Name2 string // note2
}
`
}
