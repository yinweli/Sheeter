package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

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

	bytes, err := ioutil.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, mockWriteJsonString(), string(bytes[:]))

	err = os.RemoveAll(OutputPathJson)
	assert.Nil(t, err)
}

func TestWriteJsonFailed(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	cargo := mockWriteJsonCargo()
	cargo.Columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"x", "2", "3"}},
	}
	_, err := WriteJson(cargo)
	assert.NotNil(t, err)

	err = os.RemoveAll(OutputPathJson)
	assert.Nil(t, err)
}

func mockWriteJsonCargo() *Cargo {
	return &Cargo{
		Progress: NewProgress(0, "test", ioutil.Discard),
		Element: &Element{
			Excel: "real.xlsx",
			Sheet: "data",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3"}},
			{Note: "note1", Name: "name1", Field: &FieldBool{}, Datas: []string{"false", "true", "false"}},
			{Note: "note2", Name: "name2", Field: &FieldText{}, Datas: []string{"text1", "text2", "text3"}},
		},
	}
}

func mockWriteJsonString() string {
	return `[
    {
        "name0": 1,
        "name1": false,
        "name2": "text1"
    },
    {
        "name0": 2,
        "name1": true,
        "name2": "text2"
    },
    {
        "name0": 3,
        "name1": false,
        "name2": "text3"
    }
]`
}
