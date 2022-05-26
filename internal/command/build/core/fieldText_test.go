package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldText(t *testing.T) {
	field := FieldText{}

	assert.Equal(t, "text", field.TypeExcel())
	assert.Equal(t, "std::string", field.TypeCpp())
	assert.Equal(t, "string", field.TypeCs())
	assert.Equal(t, "string", field.TypeGo())
	assert.Equal(t, false, field.Hide())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform("this is a string")
	assert.Nil(t, err)
	assert.Equal(t, "this is a string", result)
}
