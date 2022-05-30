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
	assert.Equal(t, "TestExcelTestSheet#TestSheet", string(result[:]))

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
	columns := []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
	}
	setline(columns)
	assert.Equal(t, "", newline())

	columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}},
	}
	setline(columns)
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "", newline())

	columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}},
		{Note: "note2", Name: "name2", Field: &FieldInt{}},
	}
	setline(columns)
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "\n", newline())
	assert.Equal(t, "", newline())

	columns = []*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}},
		{Note: "note2", Name: "name2", Field: &FieldInt{}},
		{Note: "note3", Name: "name3", Field: &FieldEmpty{}},
	}
	setline(columns)
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
