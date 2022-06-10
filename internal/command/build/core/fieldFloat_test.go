package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldFloat(t *testing.T) {
	field := mockFieldFloat()
	assert.Equal(t, "float", field.TypeExcel())
	assert.Equal(t, "float", field.TypeCpp())
	assert.Equal(t, "float", field.TypeCs())
	assert.Equal(t, "float32", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("0.000001")
	assert.Nil(t, err)
	assert.Equal(t, 0.000001, result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldFloat() *FieldFloat {
	return &FieldFloat{}
}
