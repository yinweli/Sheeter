package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldEmpty(t *testing.T) { // TODO: 做到這邊
	field := mockFieldEmpty()
	assert.Equal(t, "empty", field.Type())
	assert.Equal(t, false, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	assert.Equal(t, nil, field.ToJsonDefault())

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
