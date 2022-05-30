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
	assert.Equal(t, "TestExcelTestSheet#TestSheet", result.String())

	result, err = CodeGenerate("{{{.Unknown}}", cargo)
	assert.NotNil(t, err)

	result, err = CodeGenerate(code, nil)
	assert.NotNil(t, err)
}

func TestCppNamespace(t *testing.T) {
	assert.Equal(t, CppNamespace, cppNamespace())
}

func TestCsNameSpace(t *testing.T) {
	assert.Equal(t, CsNamespace, csNameSpace())
}

func TestGoPackage(t *testing.T) {
	assert.Equal(t, GoPackage, goPackage())
}

func TestStructName(t *testing.T) {
	cargo := mockCodeGenerateCargo()

	assert.Equal(t, "TestExcelTestSheet", structName(cargo))
}

func TestMemberName(t *testing.T) {
	assert.Equal(t, "TestMember", memberName("testMember"))
}

func TestNewLine(t *testing.T) {
	setline(1)
	assert.Equal(t, "\n", newline())
	assert.NotEqual(t, "\n", newline())

	setline(2)
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "", newline())

	setline(3)
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "", newline())
}

func mockCodeGenerateCargo() *Cargo {
	return &Cargo{
		Element: &Element{
			Excel: "testExcel",
			Sheet: "testSheet",
		},
	}
}
