package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldInt(t *testing.T) {
	field := FieldInt{}
	assert.Equal(t, "int", field.TypeExcel())
	assert.Equal(t, "int32_t", field.TypeCpp())
	assert.Equal(t, "int", field.TypeCs())
	assert.Equal(t, "int32", field.TypeGo())
	assert.Equal(t, true, field.Show())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform("999")
	assert.Nil(t, err)
	assert.Equal(t, int64(999), result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
