package core

import (
	"testing"

	"Sheeter/testdata"
)

func TestReadSheet(t *testing.T) {
	task := &Task{
		Global: &Global{
			ExcelPath:   testdata.RootPath,
			LineOfNote:  1,
			LineOfField: 2,
			LineOfData:  3,
		},
		Element: &Element{
			Excel: testdata.TestExcel,
			Sheet: testdata.TestSheet,
		},
	}

	ReadSheet(task)
}
