package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldPkey(t *testing.T) {
	field := mockFieldPkey()
	assert.Equal(t, "pkey", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, true, field.IsPkey())
	assert.Equal(t, int64(0), field.ToJsonDefault())

	result, err := field.ToJsonValue("123456789")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456789), result)
	_, err = field.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("123456789")
	assert.Nil(t, err)
	assert.Equal(t, "123456789", result)
	_, err = field.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(t, err)
}

func mockFieldPkey() *FieldPkey {
	return &FieldPkey{}
}
