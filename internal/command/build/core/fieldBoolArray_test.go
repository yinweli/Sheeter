package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldBoolArray(t *testing.T) {
	var result interface{}
	var err error
	var object FieldBoolArray

	assert.Equal(t, "boolArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<bool>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<bool>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]bool", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform(testdata.BoolString())
	assert.Equal(t, testdata.BoolArray(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
