package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldBoolArray(t *testing.T) {
	field := mockFieldBoolArray()
	assert.Equal(t, "boolArray", field.TypeExcel())
	assert.Equal(t, "List<bool>", field.TypeCs())
	assert.Equal(t, "[]bool", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.ToJsonValue("true,false,true,false,true")
	assert.Nil(t, err)
	assert.Equal(t, []bool{true, false, true, false, true}, result)
	result, err = field.ToJsonValue("?????")
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("true,false,true,false,true")
	assert.Nil(t, err)
	assert.Equal(t, "{true,false,true,false,true}", result)
	result, err = field.ToLuaValue("?????")
	assert.NotNil(t, err)
}

func mockFieldBoolArray() *FieldBoolArray {
	return &FieldBoolArray{}
}
