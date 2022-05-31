package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteJson(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	cargo := mockWriteJsonCargo()
	path, err := WriteJson(cargo)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(OutputPathJson, "realData.json"), path)
	assert.FileExists(t, path)

	cargo = mockWriteJsonCargo()
	cargo.Columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{"x", "2", "3", "4", "5"}},
		{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
	}
	path, err = WriteJson(cargo)
	assert.NotNil(t, err)

	err = os.RemoveAll(OutputPathJson)
	assert.Nil(t, err)
}

func mockWriteJsonCargo() *Cargo {
	return &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Element: &Element{
			Excel: "real.xlsx",
			Sheet: "data",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
		},
	}
}
