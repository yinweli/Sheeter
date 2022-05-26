package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	parser := Parser{
		fields: make(map[string]Field),
	}
	fieldBool := &FieldBool{}

	parser.Add(fieldBool)
	result, ok := parser.fields[fieldBool.TypeExcel()]
	assert.True(t, ok, "parser add failed")
	assert.NotNil(t, result, "parser add failed")
	assert.Equal(t, fieldBool.TypeExcel(), result.TypeExcel(), "parser add failed")

	name, field, err := parser.Parse("real#bool")
	assert.Nil(t, err, "parser parse real failed")
	assert.Equal(t, "real", name, "parser parse real failed")
	assert.Equal(t, (&FieldBool{}).TypeExcel(), field.TypeExcel(), "parser parse real failed")

	name, field, err = parser.Parse("fake#fake")
	assert.NotNil(t, err, "parser parse fake failed")

	name, field, err = parser.Parse("unknown")
	assert.NotNil(t, err, "parser parse unknown failed")
}

func TestNewParser(t *testing.T) {
	assert.NotNil(t, NewParser(), "new parser failed")
}
