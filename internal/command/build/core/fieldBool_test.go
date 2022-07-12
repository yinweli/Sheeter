package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldBool(t *testing.T) {
	field := mockFieldBool()
	assert.Equal(t, "bool", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	assert.Equal(t, false, field.ToJsonDefault())

	result, err := field.ToJsonValue("true")
	assert.Nil(t, err)
	assert.Equal(t, true, result)
	result, err = field.ToJsonValue("false")
	assert.Nil(t, err)
	assert.Equal(t, false, result)
	_, err = field.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("true")
	assert.Nil(t, err)
	assert.Equal(t, "true", result)
	result, err = field.ToLuaValue("false")
	assert.Nil(t, err)
	assert.Equal(t, "false", result)
	_, err = field.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(t, err)
}

func mockFieldBool() *FieldBool {
	return &FieldBool{}
}
