package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestCargo(t *testing.T) {
	cargo := mockCargo()

	assert.Equal(t, "realData.json", cargo.JsonFileName())
	assert.Equal(t, "realData.hpp", cargo.CppFileName())
	assert.Equal(t, "realData.cs", cargo.CsFileName())
	assert.Equal(t, "realData.go", cargo.GoFileName())
	assert.Equal(t, "RealData", cargo.StructName())
	assert.Equal(t, "realData.test", cargo.fileName("test"))
	assert.Equal(t, "real", cargo.excelName())
}

func TestColumn(t *testing.T) {
	cargo := mockCargo()

	assert.Equal(t, "Test", cargo.Columns[0].MemberName())
}

func mockCargo() *Cargo {
	return &Cargo{
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
		Columns: []*Column{
			{Name: "test"},
		},
	}
}
