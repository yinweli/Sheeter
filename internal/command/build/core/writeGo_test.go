package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestWriteGo(t *testing.T) {
	cargo := mockWriteGoCargo()
	path, err := WriteGo(cargo)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(OutputPathGo, "realData.go"), path)
	assert.FileExists(t, path)

	err = os.RemoveAll(OutputPathGo)
	assert.Nil(t, err)
}

func mockWriteGoCargo() *Cargo {
	return &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Element: &Element{
			Excel: "real.xlsx",
			Sheet: "data",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}},
		},
	}
}
