package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestDoubleArray(t *testing.T) {
	var result interface{}
	var err error
	var object DoubleArray

	assert.Equal(t, "doubleArray", object.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<double>", object.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<double>", object.TypeCs(), "type cs failed")
	assert.Equal(t, "[]float64", object.TypeGo(), "type go failed")
	assert.Equal(t, false, object.Hide(), "hide failed")
	assert.Equal(t, false, object.PrimaryKey(), "primary key failed")

	result, err = object.Transform(testdata.Float64String())
	assert.Equal(t, testdata.Float64Array(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = object.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
