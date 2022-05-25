package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestIntArray(t *testing.T) {
	var result interface{}
	var err error
	var object IntArray

	assert.Equal(t, "intArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<int32_t>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<int>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]int32", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform(testdata.Int32String())
	assert.Equal(t, testdata.Int32Array(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
