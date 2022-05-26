package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldDouble(t *testing.T) {
	field := FieldDouble{}

	assert.Equal(t, "double", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "double", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "double", field.TypeCs(), "type cs failed")
	assert.Equal(t, "float64", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("0.99999999")
	assert.Nil(t, err, "transform failed")
	assert.InDelta(t, 0.99999999, result, 0.00000001, "fill to metas failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
