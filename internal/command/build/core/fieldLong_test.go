package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldLong(t *testing.T) {
	field := FieldLong{}

	assert.Equal(t, "long", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "int64_t", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "long", field.TypeCs(), "type cs failed")
	assert.Equal(t, "int64", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("999999999999")
	assert.Nil(t, err, "transform failed")
	assert.Equal(t, int64(999999999999), result, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
