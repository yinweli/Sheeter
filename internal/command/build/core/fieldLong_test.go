package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldLong(t *testing.T) {
	field := mockFieldLong()
	assert.Equal(t, "long", field.TypeExcel())
	assert.Equal(t, "long", field.TypeCs())
	assert.Equal(t, "int64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

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

func mockFieldLong() *FieldLong {
	return &FieldLong{}
}
