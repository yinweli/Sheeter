package core

import (
	"testing"

	"Sheeter/internal"
	"Sheeter/testdata"

	"github.com/schollz/progressbar/v3"
)

func TestReadSheet(t *testing.T) {
	task := &Task{
		Progress: progressbar.Default(internal.ProgressDefault),
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
