package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldDoubleArray(t *testing.T) {
	field := mockFieldDoubleArray()
	assert.Equal(t, "doubleArray", field.TypeExcel())
	assert.Equal(t, "std::vector<double>", field.TypeCpp())
	assert.Equal(t, "List<double>", field.TypeCs())
	assert.Equal(t, "[]float64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("0.000001,0.000002,0.000003")
	assert.Nil(t, err)
	assert.Equal(t, []float64{0.000001, 0.000002, 0.000003}, result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldDoubleArray() *FieldDoubleArray {
	return &FieldDoubleArray{}
}
