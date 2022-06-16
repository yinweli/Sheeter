package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldFloatArray(t *testing.T) {
	field := mockFieldFloatArray()
	assert.Equal(t, "floatArray", field.TypeExcel())
	assert.Equal(t, "List<float>", field.TypeCs())
	assert.Equal(t, "[]float32", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.ToJsonValue("0.000001,0.000002,0.000003")
	assert.Nil(t, err)
	assert.Equal(t, []float64{0.000001, 0.000002, 0.000003}, result)
	result, err = field.ToJsonValue("?????")
	assert.NotNil(t, err)
}

func mockFieldFloatArray() *FieldFloatArray {
	return &FieldFloatArray{}
}
