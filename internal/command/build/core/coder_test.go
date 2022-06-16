package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoder(t *testing.T) {
	columns := []*Column{
		{Field: &FieldInt{}},
		{Field: &FieldInt{}},
		{Field: &FieldEmpty{}},
	}
	originalName := "test1"
	jsonFileName := "test2"
	structName := "test3"
	coder := NewCoder(columns, originalName, jsonFileName, structName)
	bytes, err := coder.Generate("{{.OriginalName}}#{{.JsonFileName}}#{{.StructName}}#{{.CsNamespace}}#{{.GoPackage}}#{{.FirstUpper \"testColumn\"}}")
	assert.Nil(t, err)
	assert.Equal(t, "test1#test2#test3#Sheeter#sheeter#TestColumn", string(bytes[:]))
	bytes, err = coder.Generate("{{{}}")
	assert.NotNil(t, err)
	bytes, err = coder.Generate("{{.Unknown}}")
	assert.NotNil(t, err)
	assert.Equal(t, originalName, coder.OriginalName())
	assert.Equal(t, jsonFileName, coder.JsonFileName())
	assert.Equal(t, structName, coder.StructName())
	assert.Equal(t, CsNamespace, coder.CsNamespace())
	assert.Equal(t, GoPackage, coder.GoPackage())
	assert.Equal(t, "", coder.SetLine())
	assert.Equal(t, "\n", coder.NewLine())
	assert.Equal(t, "", coder.NewLine())
	assert.Equal(t, "", coder.NewLine())
	assert.Equal(t, "TestColumn", coder.FirstUpper("testColumn"))
	assert.Equal(t, "testColumn", coder.FirstLower("TestColumn"))
}

func TestNewCoder(t *testing.T) {
	coder := NewCoder(nil, "", "", "")
	assert.NotNil(t, coder)
}

func TestCalcMaxLine(t *testing.T) {
	assert.Equal(t, 0, calcMaxLine([]*Column{
		{Field: &FieldInt{}},
		{Field: &FieldEmpty{}},
	}))
	assert.Equal(t, 1, calcMaxLine([]*Column{
		{Field: &FieldInt{}},
		{Field: &FieldInt{}},
		{Field: &FieldEmpty{}},
	}))
	assert.Equal(t, 2, calcMaxLine([]*Column{
		{Field: &FieldInt{}},
		{Field: &FieldInt{}},
		{Field: &FieldInt{}},
		{Field: &FieldEmpty{}},
	}))
	assert.Equal(t, 2, calcMaxLine([]*Column{
		{Field: &FieldInt{}},
		{Field: &FieldEmpty{}},
		{Field: &FieldInt{}},
		{Field: &FieldEmpty{}},
		{Field: &FieldInt{}},
	}))
}
