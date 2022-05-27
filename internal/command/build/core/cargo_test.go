package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestCargo(t *testing.T) {
	cargo := mockCargo()

	assert.Equal(t, "realdata.json", cargo.JsonFileName())
	assert.Equal(t, "realdata.hpp", cargo.CppFileName())
	assert.Equal(t, "realdata.cs", cargo.CsFileName())
	assert.Equal(t, "realdata.go", cargo.GoFileName())
	assert.Equal(t, "realdata.test", cargo.outputFileName("test"))
}

func mockCargo() *Cargo {
	return &Cargo{
		Element: &Element{
			Excel: testdata.RealExcel,
			Sheet: testdata.RealSheet,
		},
	}
}
