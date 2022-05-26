package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldDoubleArray(t *testing.T) {
	field := FieldDoubleArray{}

	assert.Equal(t, "doubleArray", field.TypeExcel())
	assert.Equal(t, "std::vector<double>", field.TypeCpp())
	assert.Equal(t, "List<double>", field.TypeCs())
	assert.Equal(t, "[]float64", field.TypeGo())
	assert.Equal(t, false, field.Hide())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform(testdata.Float64String())
	assert.Nil(t, err)
	assert.Equal(t, testdata.Float64Array(), result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
