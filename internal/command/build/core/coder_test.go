package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoder(t *testing.T) {
	ctx := mockCoderContext()
	result, err := Coder("{{.StructName}}#{{.CppNamespace}}", ctx)
	assert.Nil(t, err)
	assert.Equal(t, "ExcelSheet#Sheeter", string(result[:]))

	ctx = mockCoderContext()
	result, err = Coder("{{setline .Columns}}{{newline}}{{newline}}{{newline}}{{newline}}", ctx)
	assert.Nil(t, err)
	assert.Equal(t, "\n\n", string(result[:]))

	ctx = mockCoderContext()
	result, err = Coder("{{{.Unknown}}", ctx)
	assert.NotNil(t, err)

	ctx = mockCoderContext()
	result, err = Coder("{{{.Unknown}}", nil)
	assert.NotNil(t, err)
}

func mockCoderContext() *Context {
	return &Context{
		Element: &Element{
			Excel: "excel.xlsx",
			Sheet: "sheet",
		},
		Columns: []*Column{
			{Note: "note0", Name: "name0", Field: &FieldInt{}},
			{Note: "note1", Name: "name1", Field: &FieldInt{}},
			{Note: "note2", Name: "name2", Field: &FieldInt{}},
		},
	}
}
