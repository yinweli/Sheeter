package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldBool(t *testing.T) {
	field := FieldBool{}

	assert.Equal(t, "bool", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "bool", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "bool", field.TypeCs(), "type cs failed")
	assert.Equal(t, "bool", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("true")
	assert.Nil(t, err, "transform failed")
	assert.Equal(t, true, result, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
