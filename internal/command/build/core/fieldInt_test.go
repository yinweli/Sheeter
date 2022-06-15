package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldInt(t *testing.T) {
	field := mockFieldInt()
	assert.Equal(t, "int", field.TypeExcel())
	assert.Equal(t, "int", field.TypeCs())
	assert.Equal(t, "int32", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("123456789")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456789), result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldInt() *FieldInt {
	return &FieldInt{}
}
