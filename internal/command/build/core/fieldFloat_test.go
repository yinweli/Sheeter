package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldFloat(t *testing.T) {
	field := FieldFloat{}

	assert.Equal(t, "float", field.TypeExcel())
	assert.Equal(t, "float", field.TypeCpp())
	assert.Equal(t, "float", field.TypeCs())
	assert.Equal(t, "float32", field.TypeGo())
	assert.Equal(t, true, field.Show())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform("0.9999")
	assert.Nil(t, err)
	assert.InDelta(t, 0.9999, result, 0.0001)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
