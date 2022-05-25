package reader

import (
	"testing"

	"Sheeter/internal/command/build/cargo"
	"Sheeter/internal/command/build/config"
	"Sheeter/testdata"
)

func TestReadSheet(t *testing.T) {
	cargo := &cargo.Cargo{
		Global: &config.Global{
			ExcelPath:   testdata.RootPath,
			LineOfNote:  1,
			LineOfField: 2,
			LineOfData:  3,
		},
		Element: &config.Element{
			Excel: testdata.TestExcel,
			Sheet: testdata.TestSheet,
		},
	}

	ReadSheet(cargo)
}
