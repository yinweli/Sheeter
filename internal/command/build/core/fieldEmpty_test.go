package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldEmpty(t *testing.T) {
	field := FieldEmpty{}

	assert.Equal(t, "empty", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "", field.TypeCs(), "type cs failed")
	assert.Equal(t, "", field.TypeGo(), "type go failed")
	assert.Equal(t, true, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform("test")
	assert.Nil(t, result, "transform failed")
	assert.Nil(t, err, "transform failed")
}
