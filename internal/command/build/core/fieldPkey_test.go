package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldPkey(t *testing.T) {
	field := FieldPkey{}
	assert.Equal(t, "pkey", field.TypeExcel())
	assert.Equal(t, CppNamespace+"::pkey", field.TypeCpp())
	assert.Equal(t, "int", field.TypeCs())
	assert.Equal(t, "int", field.TypeGo())
	assert.Equal(t, true, field.Show())
	assert.Equal(t, true, field.PrimaryKey())

	result, err := field.Transform("999")
	assert.Nil(t, err)
	assert.Equal(t, 999, result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
