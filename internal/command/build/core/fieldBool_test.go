package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldBool(t *testing.T) {
	field := mockFieldBool()
	assert.Equal(t, "bool", field.TypeExcel())
	assert.Equal(t, "bool", field.TypeCs())
	assert.Equal(t, "bool", field.TypeGo())
	assert.Equal(t, true, field.IsShow())
	assert.Equal(t, false, field.IsPkey())
	result, err := field.Transform("true")
	assert.Nil(t, err)
	assert.Equal(t, true, result)
	result, err = field.Transform("?????")
	assert.NotNil(t, err)
}

func mockFieldBool() *FieldBool {
	return &FieldBool{}
}
