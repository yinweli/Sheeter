package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {
	var result interface{}
	var err error
	var object FieldFloat

	assert.Equal(t, "float", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "float", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "float", object.TypeCs(), "type cs failed")
	assert.Equal(t, "float32", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("0.9999")
	assert.InDelta(t, 0.9999, result, 0.0001, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
