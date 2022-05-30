package core

import (
	"io/ioutil"
	"os"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteCpp(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	cargo := mockWriteCppCargo()
	err := WriteCpp(cargo)
	assert.Nil(t, err)

	cargo = mockWriteCppCargo()
	cargo.Global = nil
	err = WriteCpp(cargo)
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
