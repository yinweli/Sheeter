package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldInt(t *testing.T) {
	field := FieldInt{}

	assert.Equal(t, "int", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "int32_t", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "int", field.TypeCs(), "type cs failed")
	assert.Equal(t, "int32", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("999")
	assert.Equal(t, int64(999), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
