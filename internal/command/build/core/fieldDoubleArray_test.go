package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldDoubleArray(t *testing.T) {
	field := mockFieldDoubleArray()
	assert.Equal(t, "doubleArray", field.TypeExcel())
	assert.Equal(t, "List<double>", field.TypeCs())
	assert.Equal(t, "[]float64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, []float64{0.123, 0.456, 0.789}, result)
	result, err = field.ToJsonValue("?????")
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, "{0.123,0.456,0.789}", result)
	result, err = field.ToLuaValue("?????")
	assert.NotNil(t, err)
}

func mockFieldDoubleArray() *FieldDoubleArray {
	return &FieldDoubleArray{}
}
