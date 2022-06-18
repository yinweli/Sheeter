package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldTextArray(t *testing.T) {
	field := mockFieldTextArray()
	assert.Equal(t, "textArray", field.Type())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.ToJsonValue("ball,book,pack")
	assert.Nil(t, err)
	assert.Equal(t, []string{"ball", "book", "pack"}, result)

	result, err = field.ToLuaValue("ball,book,pack")
	assert.Nil(t, err)
	assert.Equal(t, "{ball,book,pack}", result)
}

func mockFieldTextArray() *FieldTextArray {
	return &FieldTextArray{}
}
