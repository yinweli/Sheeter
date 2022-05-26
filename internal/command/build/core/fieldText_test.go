package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldText(t *testing.T) {
	var result interface{}
	var err error
	var object FieldText

	assert.Equal(t, "text", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::string", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "string", object.TypeCs(), "type cs failed")
	assert.Equal(t, "string", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform("this is a string")
	assert.Equal(t, "this is a string", result, "transform failed")
	assert.Nil(t, err, "transform failed")
}
