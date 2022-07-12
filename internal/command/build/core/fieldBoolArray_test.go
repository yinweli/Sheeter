package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldBoolArray(t *testing.T) {
	field := mockFieldBoolArray()
	assert.Equal(t, "boolArray", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	assert.Equal(t, []bool{}, field.ToJsonDefault())

	result, err := field.ToJsonValue("true,false,true,false,true")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, result)
	_, err = field.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("true,false,true,false,true")
	assert.Nil(t, err)
	assert.Equal(t, "{true,false,true,false,true}", result)
	_, err = field.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(t, err)
}

func mockFieldBoolArray() *FieldBoolArray {
	return &FieldBoolArray{}
}
