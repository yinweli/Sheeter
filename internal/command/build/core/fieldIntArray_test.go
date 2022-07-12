package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldIntArray(t *testing.T) {
	field := mockFieldIntArray()
	assert.Equal(t, "intArray", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	assert.Equal(t, []int64{}, field.ToJsonDefault())

	result, err := field.ToJsonValue("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, []int64{123, 456, 789}, result)
	_, err = field.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("123,456,789")
	assert.Nil(t, err)
	assert.Equal(t, "{123,456,789}", result)
	_, err = field.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(t, err)
}

func mockFieldIntArray() *FieldIntArray {
	return &FieldIntArray{}
}
