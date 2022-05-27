package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/internal/util"
	"Sheeter/testdata"
)

func TestWriteFile(t *testing.T) {
	// cargo := mockWriteJsonCargo()
	// bytes := []byte(testdata.Text)
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
