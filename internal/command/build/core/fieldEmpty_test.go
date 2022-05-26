package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldEmpty(t *testing.T) {
	field := FieldEmpty{}

	assert.Equal(t, "empty", field.TypeExcel())
	assert.Equal(t, "", field.TypeCpp())
	assert.Equal(t, "", field.TypeCs())
	assert.Equal(t, "", field.TypeGo())
	assert.Equal(t, true, field.Hide())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform("test")
	assert.Nil(t, err)
	assert.Nil(t, result)
}
