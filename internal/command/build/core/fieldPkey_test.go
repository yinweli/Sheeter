package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldPkey(t *testing.T) {
	field := mockFieldPkey()
	assert.Equal(t, "pkey", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, true, field.IsPkey())

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

func mockFieldPkey() *FieldPkey {
	return &FieldPkey{}
}
