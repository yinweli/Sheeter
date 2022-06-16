package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldLongArray(t *testing.T) {
	field := mockFieldLongArray()
	assert.Equal(t, "longArray", field.TypeExcel())
	assert.Equal(t, "List<long>", field.TypeCs())
	assert.Equal(t, "[]int64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.ToJsonValue("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, []int64{123, 456, 789}, result)
	result, err = field.ToJsonValue("?????")
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, "{123,456,789}", result)
	result, err = field.ToLuaValue("?????")
	assert.NotNil(t, err)
}

func mockFieldLongArray() *FieldLongArray {
	return &FieldLongArray{}
}
