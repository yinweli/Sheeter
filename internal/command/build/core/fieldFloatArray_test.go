package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldFloatArray(t *testing.T) {
	field := mockFieldFloatArray()
	assert.Equal(t, "floatArray", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	assert.Equal(t, []float64{}, field.ToJsonDefault())

	result, err := field.ToJsonValue("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, []float64{0.123, 0.456, 0.789}, result)
	_, err = field.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("0.123,0.456,0.789")
	assert.Nil(t, err)
	assert.Equal(t, "{0.123,0.456,0.789}", result)
	_, err = field.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(t, err)
}

func mockFieldFloatArray() *FieldFloatArray {
	return &FieldFloatArray{}
}
