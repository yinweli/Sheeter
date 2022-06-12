package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldEmpty(t *testing.T) {
	field := mockFieldEmpty()
	assert.Equal(t, "empty", field.TypeExcel())
	assert.Equal(t, "", field.TypeCpp())
	assert.Equal(t, "", field.TypeCs())
	assert.Equal(t, "", field.TypeGo())
	assert.Equal(t, false, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("test")
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func mockFieldEmpty() *FieldEmpty {
	return &FieldEmpty{}
}
