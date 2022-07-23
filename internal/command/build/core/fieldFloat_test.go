package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yinweli/Sheeter/testdata"
)

func TestFieldFloat(t *testing.T) { // TODO: 做到這邊
	field := mockFieldFloat()
	assert.Equal(t, "float", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	assert.Equal(t, float64(0), field.ToJsonDefault())

	result, err := field.ToJsonValue("0.123456")
	assert.Nil(t, err)
	assert.Equal(t, 0.123456, result)
	_, err = field.ToJsonValue(testdata.UnknownStr)
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("0.123456")
	assert.Nil(t, err)
	assert.Equal(t, "0.123456", result)
	_, err = field.ToLuaValue(testdata.UnknownStr)
	assert.NotNil(t, err)
}

func mockFieldFloat() *FieldFloat {
	return &FieldFloat{}
}
