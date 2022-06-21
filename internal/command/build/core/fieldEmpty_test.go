package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldEmpty(t *testing.T) {
	field := mockFieldEmpty()
	assert.Equal(t, "empty", field.Type())
	assert.Equal(t, false, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.ToJsonValue("test")
	assert.Nil(t, err)
	assert.Nil(t, result)

	result, err = field.ToLuaValue("test")
	assert.Nil(t, err)
	assert.Equal(t, "", result)
}

func mockFieldEmpty() *FieldEmpty {
	return &FieldEmpty{}
}
