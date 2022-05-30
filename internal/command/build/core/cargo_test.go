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
	assert.Equal(t, "real", cargo.PureExcelName())
	assert.Equal(t, "realData.test", cargo.outputFileName("test"))
}

func mockCargo() *Cargo {
	return &Cargo{
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
	}
}
