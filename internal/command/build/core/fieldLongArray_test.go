package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldLongArray(t *testing.T) {
	var result interface{}
	var err error
	var object FieldLongArray

	assert.Equal(t, "longArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<int64_t>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<long>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]int64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform(testdata.Int64String())
	assert.Equal(t, testdata.Int64Array(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
