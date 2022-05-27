package core

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	cargo := mockWriteJsonCargo()
	input := []byte(testdata.Text)
	filePath := filepath.Join(OutputPathJson, cargo.JsonFileName())

	assert.Nil(t, writeFile(cargo, input))

	output, _ := ioutil.ReadFile(filePath)

	assert.Equal(t, input, output)
}

func mockWriteJsonCargo() *Cargo {
	return &Cargo{
		Progress: util.NewProgressBar("test", ioutil.Discard),
		Global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfNote:  1,
			LineOfField: 2,
			LineOfData:  3,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}, Datas: []string{"1", "2", "3", "4", "5"}},
		},
	}
}
