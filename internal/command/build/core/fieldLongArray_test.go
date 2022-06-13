package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldLongArray(t *testing.T) {
	field := mockFieldLongArray()
	assert.Equal(t, "longArray", field.TypeExcel())
	assert.Equal(t, "std::vector<int64_t>", field.TypeCpp())
	assert.Equal(t, "List<long>", field.TypeCs())
	assert.Equal(t, "[]int64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, []int64{123, 456, 789}, result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldLongArray() *FieldLongArray {
	return &FieldLongArray{}
}
