package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
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
	_, err = field.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("123456789")
	assert.Nil(t, err)
	assert.Equal(t, "123456789", result)
	_, err = field.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(t, err)
}

func mockFieldInt() *FieldInt {
	return &FieldInt{}
}
