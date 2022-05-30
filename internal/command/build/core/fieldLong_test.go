package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldLong(t *testing.T) {
	field := FieldLong{}

	assert.Equal(t, "long", field.TypeExcel())
	assert.Equal(t, "int64_t", field.TypeCpp())
	assert.Equal(t, "long", field.TypeCs())
	assert.Equal(t, "int64", field.TypeGo())
	assert.Equal(t, true, field.Show())
	assert.Equal(t, false, field.PrimaryKey())

	result, err := field.Transform("999999999999")
	assert.Nil(t, err)
	assert.Equal(t, int64(999999999999), result)

	result, err = field.Transform("fake")
	assert.NotNil(t, err)
}
