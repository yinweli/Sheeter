package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldLong(t *testing.T) {
	field := mockFieldLong()
	assert.Equal(t, "long", field.TypeExcel())
	assert.Equal(t, "int64_t", field.TypeCpp())
	assert.Equal(t, "long", field.TypeCs())
	assert.Equal(t, "int64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("123456789")
	assert.Nil(t, err)
	assert.Equal(t, int64(123456789), result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldLong() *FieldLong {
	return &FieldLong{}
}
