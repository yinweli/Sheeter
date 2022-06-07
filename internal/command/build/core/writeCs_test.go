package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteCs(t *testing.T) {
	dir := testdata.ChangeWorkDir()
	defer testdata.RestoreWorkDir(dir)

	writeCs := mockWriteCs()
	assert.Equal(t, "cs", writeCs.Long())
	assert.Equal(t, "s", writeCs.Short())
	assert.Equal(t, "generate cs file", writeCs.Note())
	assert.Equal(t, 2, writeCs.Calc(0))

	cargo := mockWriteCsCargo()
	path, err := writeCs.Execute(cargo)
	assert.Nil(t, err)
	assert.Equal(t, filepath.Join(OutputPathCs, "realData.cs"), path)
	assert.FileExists(t, path)

	bytes, err := ioutil.ReadFile(path)
	assert.Nil(t, err)
	assert.Equal(t, mockWriteCsString(), string(bytes[:]))

	err = os.RemoveAll(OutputPathCs)
	assert.Nil(t, err)
}

func mockWriteCs() *WriteCs {
	return &WriteCs{}
}

func mockWriteCsCargo() *Cargo {
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

func mockWriteCsString() string {
	return `// generation by sheeter ^o<

using System;
using System.Collections.Generic;

namespace Sheeter {
    public class RealData { 
        public const string fileName = "realData.json";
        public int Name0; // note0
        public bool Name1; // note1
        public string Name2; // note2
    }
} // namespace Sheeter`
}
