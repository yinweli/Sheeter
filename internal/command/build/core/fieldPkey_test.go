package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldPkey(t *testing.T) {
	field := FieldPkey{}

	assert.Equal(t, "pkey", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "Sheet::pkey", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", field.TypeCs(), "type cs failed")
	assert.Equal(t, "int", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, true, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("999")
	assert.Nil(t, err, "transform failed")
	assert.Equal(t, 999, result, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
