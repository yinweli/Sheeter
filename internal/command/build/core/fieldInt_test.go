package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldInt(t *testing.T) {
	field := mockFieldInt()
	assert.Equal(t, "int", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	assert.Equal(t, int64(0), field.ToJsonDefault())

	result, err := field.ToJsonValue("123456789")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456789), result)
	result, err = field.ToJsonValue("?????")
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("123456789")
	assert.Nil(t, err)
	assert.Equal(t, "123456789", result)
	result, err = field.ToLuaValue("?????")
	assert.NotNil(t, err)
}

func mockFieldInt() *FieldInt {
	return &FieldInt{}
}
