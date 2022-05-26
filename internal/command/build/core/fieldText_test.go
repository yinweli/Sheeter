package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldText(t *testing.T) {
	field := FieldText{}

	assert.Equal(t, "text", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::string", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "string", field.TypeCs(), "type cs failed")
	assert.Equal(t, "string", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("this is a string")
	assert.Equal(t, "this is a string", result, "transform failed")
	assert.Nil(t, err, "transform failed")
}
