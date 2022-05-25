package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	var result interface{}
	var err error
	var object Bool

	assert.Equal(t, "bool", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "bool", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "bool", object.TypeCs(), "type cs failed")
	assert.Equal(t, "bool", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("true")
	assert.Equal(t, true, result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
