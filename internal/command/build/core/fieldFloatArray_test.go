package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldFloatArray(t *testing.T) {
	field := mockFieldFloatArray()
	assert.Equal(t, "floatArray", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, []float64{0.123, 0.456, 0.789}, result)
	result, err = field.ToJsonValue("?????")
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, "{0.123,0.456,0.789}", result)
	result, err = field.ToLuaValue("?????")
	assert.NotNil(t, err)
}

func mockFieldFloatArray() *FieldFloatArray {
	return &FieldFloatArray{}
}
