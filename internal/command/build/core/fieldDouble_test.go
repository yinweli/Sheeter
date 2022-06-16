package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldDouble(t *testing.T) {
	field := mockFieldDouble()
	assert.Equal(t, "double", field.TypeExcel())
	assert.Equal(t, "double", field.TypeCs())
	assert.Equal(t, "float64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.ToJsonValue("0.123456")
	assert.Nil(t, err)
	assert.Equal(t, 0.123456, result)
	result, err = field.ToJsonValue("?????")
	assert.NotNil(t, err)

	result, err = field.ToLuaValue("0.123456")
	assert.Nil(t, err)
	assert.Equal(t, "0.123456", result)
	result, err = field.ToLuaValue("?????")
	assert.NotNil(t, err)
}

func mockFieldDouble() *FieldDouble {
	return &FieldDouble{}
}
