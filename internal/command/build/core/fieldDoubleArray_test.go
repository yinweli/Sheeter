package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldDoubleArray(t *testing.T) {
	field := FieldDoubleArray{}

	assert.Equal(t, "doubleArray", field.TypeExcel(), "type excel failed")
	assert.Equal(t, "std::vector<double>", field.TypeCpp(), "type cpp failed")
	assert.Equal(t, "List<double>", field.TypeCs(), "type cs failed")
	assert.Equal(t, "[]float64", field.TypeGo(), "type go failed")
	assert.Equal(t, false, field.Hide(), "hide failed")
	assert.Equal(t, false, field.PrimaryKey(), "primary key failed")

	result, err := field.Transform(testdata.Float64String())
	assert.Equal(t, testdata.Float64Array(), result, "transform failed")
	assert.Nil(t, err, "transform failed")

	result, err = field.Transform("fake")
	assert.NotNil(t, err, "transform failed")
}
