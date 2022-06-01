package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldDouble(t *testing.T) {
	field := FieldDouble{}
	assert.Equal(t, "double", field.TypeExcel())
	assert.Equal(t, "double", field.TypeCpp())
	assert.Equal(t, "double", field.TypeCs())
	assert.Equal(t, "float64", field.TypeGo())
	assert.Equal(t, true, field.Show())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform("0.99999999")
	assert.Nil(t, err)
	assert.InDelta(t, 0.99999999, result, 0.00000001)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
