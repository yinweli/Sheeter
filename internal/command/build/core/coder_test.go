package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoder(t *testing.T) {
	coder := mockCoder()
	coder.Template = "{{cppNamespace}}#{{.StructName}}"
	result, err := coder.Execute()
	assert.Nil(t, err)
	assert.Equal(t, "Sheeter#TestExcelTestSheet", string(result[:]))

	coder = mockCoder()
	coder.Template = "{{{.Unknown}}"
	result, err = coder.Execute()
	assert.NotNil(t, err)

	coder = mockCoder()
	coder.Template = "{{cppNamespace}}#{{.StructName}}"
	coder.Cargo = nil
	result, err = coder.Execute()
	assert.NotNil(t, err)

	coder = mockCoder()
	assert.Equal(t, CppNamespace, coder.cppNamespace())
	assert.Equal(t, CsNamespace, coder.csNamespace())
	assert.Equal(t, GoPackage, coder.goPackage())

	coder = mockCoder()
	coder.setline([]*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
	})
	assert.Equal(t, "", coder.newline())

	coder = mockCoder()
	coder.setline([]*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}},
	})
	assert.Equal(t, "\n", coder.newline())
	assert.Equal(t, "", coder.newline())

	coder = mockCoder()
	coder.setline([]*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}},
		{Note: "note2", Name: "name2", Field: &FieldInt{}},
	})
	assert.Equal(t, "\n", coder.newline())
	assert.Equal(t, "\n", coder.newline())
	assert.Equal(t, "", coder.newline())

	coder = mockCoder()
	coder.setline([]*Column{
		{Note: "note0", Name: "name0", Field: &FieldInt{}},
		{Note: "note1", Name: "name1", Field: &FieldInt{}},
		{Note: "note2", Name: "name2", Field: &FieldInt{}},
		{Note: "note3", Name: "name3", Field: &FieldEmpty{}},
	})
	assert.Equal(t, "\n", coder.newline())
	assert.Equal(t, "\n", coder.newline())
	assert.Equal(t, "", coder.newline())
}

func TestNewCoder(t *testing.T) {
	coder := NewCoder("", nil)
	assert.NotNil(t, coder)
}

func mockCoder() *Coder {
	return &Coder{
		Cargo: &Cargo{
			Element: &Element{
				Excel: "testExcel",
				Sheet: "testSheet",
			},
		},
	}
}
