package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldFloat(t *testing.T) {
	field := FieldFloat{}

	assert.Equal(t, "float", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "float", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "float", field.TypeCs(), "type cs failed")
	assert.Equal(t, "float32", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("0.9999")
	assert.InDelta(t, 0.9999, result, 0.0001, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
