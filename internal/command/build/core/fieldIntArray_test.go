package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldIntArray(t *testing.T) {
	field := mockFieldIntArray()
	assert.Equal(t, "intArray", field.TypeExcel())
	assert.Equal(t, "List<int>", field.TypeCs())
	assert.Equal(t, "[]int32", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, []int64{123, 456, 789}, result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldIntArray() *FieldIntArray {
	return &FieldIntArray{}
}
