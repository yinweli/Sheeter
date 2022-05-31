package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestWriteCs(t *testing.T) {
	cargo := mockWriteCsCargo()
	path, err := WriteCs(cargo)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(OutputPathCs, "realData.cs"), path)
	assert.FileExists(t, path)

	err = os.RemoveAll(OutputPathCs)
	assert.Nil(t, err)
}

func mockWriteCsCargo() *Cargo {
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
