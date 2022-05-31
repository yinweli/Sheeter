package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/internal/util"

	"github.com/stretchr/testify/assert"
)

func TestWriteCpp(t *testing.T) {
	cargo := mockWriteCppCargo()
	path, err := WriteCpp(cargo)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(OutputPathCpp, "realData.hpp"), path)
	assert.FileExists(t, path)

	cargo = mockWriteCppCargo()
	cargo.Global = nil
	path, err = WriteCpp(cargo)
	assert.NotNil(t, err)

	err = os.RemoveAll(OutputPathCpp)
	assert.Nil(t, err)
}

func mockWriteCppCargo() *Cargo {
	return &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Global: &Global{
			CppLibraryPath: "nlohmann/json.hpp",
		},
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
