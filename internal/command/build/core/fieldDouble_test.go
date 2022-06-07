package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldDouble(t *testing.T) {
	field := mockFieldDouble()
	assert.Equal(t, "double", field.TypeExcel())
	assert.Equal(t, "double", field.TypeCpp())
	assert.Equal(t, "double", field.TypeCs())
	assert.Equal(t, "float64", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())

	result, err := field.Transform("0.99999999")
	assert.Nil(t, err)
	assert.InDelta(t, 0.99999999, result, 0.00000001)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}

func mockFieldDouble() *FieldDouble {
	return &FieldDouble{}
}
