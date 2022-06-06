package core

import (
	"testing"

	"Sheeter/testdata"

	"github.com/stretchr/testify/assert"
)

func TestFieldIntArray(t *testing.T) {
	field := FieldIntArray{}
	assert.Equal(t, "intArray", field.TypeExcel())
	assert.Equal(t, "std::vector<int32_t>", field.TypeCpp())
	assert.Equal(t, "List<int>", field.TypeCs())
	assert.Equal(t, "[]int32", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.Transform(testdata.Int32String())
	assert.Nil(t, err)
	assert.Equal(t, testdata.Int32Array(), result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
