package core

import (
	"io/ioutil"
	"os"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteCs(t *testing.T) {
	cargo := mockWriteCsCargo()
	filePath, err := WriteCs(cargo)
	assert.Nil(t, err)
	assert.FileExists(t, filePath)

	err = os.RemoveAll(OutputPathCs)
	assert.Nil(t, err)
}

func mockWriteCsCargo() *Cargo {
	return &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}},
		},
	}
}
