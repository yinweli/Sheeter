package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldFloatArray(t *testing.T) {
	field := mockFieldFloatArray()
	assert.Equal(t, "floatArray", field.TypeExcel())
	assert.Equal(t, "std::vector<float>", field.TypeCpp())
	assert.Equal(t, "List<float>", field.TypeCs())
	assert.Equal(t, "[]float32", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.Transform(testdata.Float32String())
	assert.Nil(t, err)
	assert.Equal(t, testdata.Float32Array(), result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}

func mockFieldFloatArray() *FieldFloatArray {
	return &FieldFloatArray{}
}
