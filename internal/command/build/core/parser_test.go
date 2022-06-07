package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	parser := mockParser()

	parser.Add(&FieldBool{})
	field, ok := parser.fields[(&FieldBool{}).TypeExcel()]
	assert.True(t, ok)
	assert.NotNil(t, field)

	name, field, err := parser.Parse("real#bool")
	assert.Nil(t, err)
	assert.Equal(t, "real", name)
	assert.Equal(t, (&FieldBool{}).TypeExcel(), field.TypeExcel())

	name, field, err = parser.Parse("fake#fake")
	assert.NotNil(t, err)

	name, field, err = parser.Parse("unknown")
	assert.NotNil(t, err)
}

func TestNewParser(t *testing.T) {
	parser := NewParser()
	assert.NotNil(t, parser)
}

func mockParser() *Parser {
	return &Parser{
		fields: map[string]Field{},
	}
}
