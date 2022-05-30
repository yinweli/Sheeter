package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeGenerate(t *testing.T) {
	code := "{{structName .}}#{{memberName .Element.Sheet}}"
	cargo := mockCodeGenerateCargo()

	result, err := CodeGenerate(code, cargo)
	assert.Nil(t, err)
	assert.Equal(t, "TestExcelTestSheet#TestSheet", result)

	result, err = CodeGenerate("{{{.Unknown}}", cargo)
	assert.NotNil(t, err)

	result, err = CodeGenerate(code, nil)
	assert.NotNil(t, err)
}

func TestStructName(t *testing.T) {
	cargo := mockCodeGenerateCargo()

	assert.Equal(t, "TestExcelTestSheet", structName(cargo))
}

func TestMemberName(t *testing.T) {
	assert.Equal(t, "TestMember", memberName("testMember"))
}

func mockCodeGenerateCargo() *Cargo {
	return &Cargo{
		Element: &Element{
			Excel: "testExcel",
			Sheet: "testSheet",
		},
	}
}
