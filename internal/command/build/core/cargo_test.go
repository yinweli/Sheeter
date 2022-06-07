package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCargo(t *testing.T) {
	cargo := mockCargo()
	assert.Equal(t, "excel.xlsx(sheet)", cargo.LogName())
	assert.Equal(t, "ExcelSheet", cargo.StructName())
	assert.Equal(t, "excelSheet.json", cargo.JsonFileName())
	assert.Equal(t, "excelSheet.hpp", cargo.CppFileName())
	assert.Equal(t, "excelSheet.cs", cargo.CsFileName())
	assert.Equal(t, "excelSheet.go", cargo.GoFileName())
	assert.Equal(t, "excelSheet.test", cargo.fileName("test"))
	assert.Equal(t, "excel", cargo.excelName())
}

func TestSheets(t *testing.T) {
	sheets := Sheets{
		{"1", "2", "3", "4", "5"},
		{"1", "2", "3", "4"},
		{"1", "2", "3"},
	}
	count := 0

	for _, itor := range sheets {
		count = count + len(itor)
	} // for

	assert.Equal(t, count, sheets.Size())
}

func TestColumn(t *testing.T) {
	cargo := mockCargo()
	assert.Equal(t, "Test", cargo.Columns[0].MemberName())
}

func mockCargo() *Cargo {
	return &Cargo{
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Columns: []*Column{
			{Name: "test"},
		},
	}
}
