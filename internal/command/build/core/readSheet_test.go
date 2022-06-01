package core

import (
	"io/ioutil"
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestReadSheet(t *testing.T) {
	cargo := mockReadSheetCargo()
	err := ReadSheet(cargo, 0, ioutil.Discard)
	assert.Nil(t, err)
	assert.NotNil(t, cargo.Progress)
	assert.NotNil(t, cargo.Sheets)
	assert.Equal(t, 12, len(cargo.Sheets))
	assert.Equal(t, 16, len(cargo.Sheets[0]))
	assert.Equal(t, "checkpoint", cargo.Sheets[0][15])
	assert.Equal(t, "checkpoint", cargo.Sheets[11][15])

	cargo = mockReadSheetCargo()
	cargo.Global.ExcelPath = ""
	err = ReadSheet(cargo, 0, ioutil.Discard)
	assert.NotNil(t, err)

	cargo = mockReadSheetCargo()
	cargo.Element.Excel = testdata.DefectExcel
	cargo.Element.Sheet = testdata.DefectSheet
	err = ReadSheet(cargo, 0, ioutil.Discard)
	assert.NotNil(t, err)

	cargo = mockReadSheetCargo()
	cargo.Element.Excel = testdata.UnknownExcel
	err = ReadSheet(cargo, 0, ioutil.Discard)
	assert.NotNil(t, err)

	cargo = mockReadSheetCargo()
	cargo.Element.Sheet = testdata.UnknownSheet
	err = ReadSheet(cargo, 0, ioutil.Discard)
	assert.NotNil(t, err)
}

func mockReadSheetCargo() *Cargo {
	return &Cargo{
		Global: &Global{
			ExcelPath: testdata.RootPath,
		},
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
	}
}
